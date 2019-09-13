package wallet

import (
	"context"
	"crypto/sha256"
	"database/sql"

	"github.com/cachecashproject/go-cachecash/ccmsg"
	"github.com/cachecashproject/go-cachecash/keypair"
	"github.com/cachecashproject/go-cachecash/ledger"
	"github.com/cachecashproject/go-cachecash/ledger/txscript"
	"github.com/cachecashproject/go-cachecash/wallet/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/ed25519"
)

type Account struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey // May be nil
}

func GenerateAccount() (*Account, error) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate keypair")
	}

	return &Account{
		PublicKey:  pub,
		PrivateKey: priv,
	}, nil
}

func (ac *Account) P2WPKHAddress(v ledger.AddressVersion) *ledger.P2WPKHAddress {
	pkh := sha256.Sum256(ac.PublicKey)

	return &ledger.P2WPKHAddress{
		AddressVersion:        v,
		WitnessProgramVersion: 0,
		PublicKeyHash:         pkh[:ledger.AddressHashSize],
	}
}

type Wallet struct {
	l    *logrus.Logger
	kp   *keypair.KeyPair
	db   *sql.DB
	grpc ccmsg.LedgerClient
}

func NewWallet(l *logrus.Logger, kp *keypair.KeyPair, db *sql.DB, grpc ccmsg.LedgerClient) *Wallet {
	return &Wallet{
		l:    l,
		kp:   kp,
		db:   db,
		grpc: grpc,
	}
}

func (w *Wallet) BlockHeight(ctx context.Context) (uint64, error) {
	count, err := models.Blocks().Count(ctx, w.db)
	return uint64(count), err
}

func (w *Wallet) FetchBlocks(ctx context.Context) error {
	height, err := w.BlockHeight(ctx)
	if err != nil {
		return err
	}

	w.l.WithFields(logrus.Fields{
		"height": height,
	}).Info("Fetching blocks")
	resp, err := w.grpc.GetBlocks(ctx, &ccmsg.GetBlocksRequest{
		StartDepth: height,
		Limit:      5,
	})
	if err != nil {
		return errors.Wrap(err, "failed to fetch blocks")
	}

	if len(resp.Blocks) == 0 {
		w.l.Info("No new blocks")
	}

	for _, bytes := range resp.Blocks {
		block := ledger.Block{}
		err = block.Unmarshal(bytes)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal block")
		}

		w.l.Info("Adding block")
		err = w.AddBlock(ctx, block)
		if err != nil {
			return err
		}

		blockModel := &models.Block{
			Height: int64(height),
			Bytes:  string(bytes),
		}
		err = blockModel.Insert(ctx, w.db, boil.Infer())
		if err != nil {
			return nil
		}
	}

	return nil
}

func (w *Wallet) MatchesOurWallet(output ledger.TransactionOutput) (bool, error) {
	pubKeyHash := txscript.Hash160Sum(w.kp.PublicKey)
	script, err := txscript.MakeP2WPKHInputScript(pubKeyHash)
	if err != nil {
		return false, errors.Wrap(err, "failed to crate p2wpkh input script")
	}

	scriptBytes, err := script.Marshal()
	if err != nil {
		return false, errors.Wrap(err, "failed to marshal input script")
	}

	matches := string(output.ScriptPubKey) == string(scriptBytes)
	return matches, nil
}

func (w *Wallet) AddBlock(ctx context.Context, block ledger.Block) error {
	for _, tx := range block.Transactions {
		txid, err := tx.TXID()
		if err != nil {
			return err
		}

		// TODO: mark outputs as spent
		for _, txo := range tx.Inputs() {
			err = w.UnspendableUTXO(ctx, txo.Outpoint)
			if err != nil {
				return errors.Wrap(err, "failed to mark utxo as spent")
			}
		}

		for idx, output := range tx.Outputs() {
			if matches, err := w.MatchesOurWallet(output); err != nil || !matches {
				continue
			}

			w.l.Info("discovered spendable transaction output")
			err = w.AddUTXO(ctx, &models.Utxo{
				Txid:         string(txid[:]),
				Idx:          int64(idx),
				Amount:       int64(output.Value),
				ScriptPubkey: string(output.ScriptPubKey),
			})
			if err != nil {
				return errors.Wrap(err, "failed to add utxo to db")
			}
		}
	}

	return nil
}

func (w *Wallet) AddUTXO(ctx context.Context, utxo *models.Utxo) error {
	return utxo.Insert(ctx, w.db, boil.Infer())
}

func (w *Wallet) GetUTXOs(ctx context.Context) ([]*models.Utxo, error) {
	return models.Utxos().All(ctx, w.db)
}

// we've spent the utxo
// TODO: refactor this
func (w *Wallet) DeleteUTXO(ctx context.Context, utxo *models.Utxo) error {
	_, err := utxo.Delete(ctx, w.db)
	return err
}

// we've observed the utxo has been spent
// TODO: refactor this
func (w *Wallet) UnspendableUTXO(ctx context.Context, utxo ledger.Outpoint) error {
	_, err := models.Utxos(qm.Where("txid = ? and idx = ?", string(utxo.PreviousTx[:]), utxo.Index)).DeleteAll(ctx, w.db)
	return err
}

func (w *Wallet) GenerateTX(ctx context.Context, target ed25519.PublicKey, amount uint32) (*ledger.Transaction, error) {
	utxos, err := w.GetUTXOs(ctx)
	if err != nil {
		return nil, err
	}

	inputs := []ledger.TransactionInput{}
	prevOutputs := []ledger.TransactionOutput{}
	outputs := []ledger.TransactionOutput{}

	spendingSum := uint32(0)
	for _, utxo := range utxos {
		spendingSum += uint32(utxo.Amount)

		txid := ledger.TXID{}
		copy(txid[:], utxo.Txid)

		pubKeyHash := txscript.Hash160Sum(w.kp.PublicKey)
		scriptSig, err := txscript.MakeP2WPKHOutputScript(pubKeyHash)
		if err != nil {
			return nil, errors.Wrap(err, "todo")
		}
		scriptSigBytes, err := scriptSig.Marshal()
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal input script")
		}

		inputs = append(inputs, ledger.TransactionInput{
			Outpoint: ledger.Outpoint{
				PreviousTx: txid,
				Index:      uint8(utxo.Idx),
			},
			ScriptSig:  scriptSigBytes,
			SequenceNo: 0xFFFFFFFF,
		})

		prevOutputs = append(prevOutputs, ledger.TransactionOutput{
			Value:        uint32(utxo.Amount),
			ScriptPubKey: []byte(utxo.ScriptPubkey),
		})

		err = w.DeleteUTXO(ctx, utxo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete spent utxo")
		}

		if spendingSum >= amount {
			break
		}
	}

	if spendingSum < amount {
		return nil, errors.New("insufficient funds")
	}

	pubKeyHash := txscript.Hash160Sum(target)
	scriptPubkey, err := txscript.MakeP2WPKHInputScript(pubKeyHash)
	if err != nil {
		return nil, errors.Wrap(err, "todo")
	}
	scriptPubkeyBytes, err := scriptPubkey.Marshal()
	if err != nil {
		return nil, errors.Wrap(err, "todo")
	}

	outputs = append(outputs, ledger.TransactionOutput{
		Value:        amount,
		ScriptPubKey: scriptPubkeyBytes,
	})

	change := spendingSum - amount
	if change > 0 {
		w.l.Info("adding change to tx: ", change)

		pubKeyHash := txscript.Hash160Sum(w.kp.PublicKey)
		scriptPubkey, err := txscript.MakeP2WPKHInputScript(pubKeyHash)
		if err != nil {
			return nil, errors.Wrap(err, "todo")
		}
		scriptPubkeyBytes, err := scriptPubkey.Marshal()
		if err != nil {
			return nil, errors.Wrap(err, "todo")
		}

		outputs = append(outputs, ledger.TransactionOutput{
			Value:        change,
			ScriptPubKey: scriptPubkeyBytes,
		})
	}

	tx := ledger.Transaction{
		Version: 1,
		Flags:   0,
		Body: &ledger.TransferTransaction{
			Inputs:   inputs,
			Outputs:  outputs,
			LockTime: 0,
		},
	}

	err = tx.GenerateWitnesses(w.kp, prevOutputs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate witnesses")
	}

	return &tx, nil
}

func (w *Wallet) SendCoins(ctx context.Context, target ed25519.PublicKey, amount uint32) error {
	w.l.Info("generating transaction")
	tx, err := w.GenerateTX(ctx, target, amount)
	if err != nil {
		return errors.Wrap(err, "failed to generate tx")
	}

	w.l.Info("sending transaction to ledgerd...")
	_, err = w.grpc.PostTransaction(ctx, &ccmsg.PostTransactionRequest{Tx: *tx})
	if err != nil {
		return errors.Wrap(err, "failed to post transaction")
	}
	w.l.Info("tx got accepted")

	return nil
}
