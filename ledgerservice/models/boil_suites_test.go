// +build sqlboiler_test

// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Blocks", testBlocks)
	t.Run("MempoolTransactions", testMempoolTransactions)
	t.Run("TransactionAuditlogs", testTransactionAuditlogs)
	t.Run("Utxos", testUtxos)
}

func TestDelete(t *testing.T) {
	t.Run("Blocks", testBlocksDelete)
	t.Run("MempoolTransactions", testMempoolTransactionsDelete)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsDelete)
	t.Run("Utxos", testUtxosDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Blocks", testBlocksQueryDeleteAll)
	t.Run("MempoolTransactions", testMempoolTransactionsQueryDeleteAll)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsQueryDeleteAll)
	t.Run("Utxos", testUtxosQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Blocks", testBlocksSliceDeleteAll)
	t.Run("MempoolTransactions", testMempoolTransactionsSliceDeleteAll)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsSliceDeleteAll)
	t.Run("Utxos", testUtxosSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Blocks", testBlocksExists)
	t.Run("MempoolTransactions", testMempoolTransactionsExists)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsExists)
	t.Run("Utxos", testUtxosExists)
}

func TestFind(t *testing.T) {
	t.Run("Blocks", testBlocksFind)
	t.Run("MempoolTransactions", testMempoolTransactionsFind)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsFind)
	t.Run("Utxos", testUtxosFind)
}

func TestBind(t *testing.T) {
	t.Run("Blocks", testBlocksBind)
	t.Run("MempoolTransactions", testMempoolTransactionsBind)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsBind)
	t.Run("Utxos", testUtxosBind)
}

func TestOne(t *testing.T) {
	t.Run("Blocks", testBlocksOne)
	t.Run("MempoolTransactions", testMempoolTransactionsOne)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsOne)
	t.Run("Utxos", testUtxosOne)
}

func TestAll(t *testing.T) {
	t.Run("Blocks", testBlocksAll)
	t.Run("MempoolTransactions", testMempoolTransactionsAll)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsAll)
	t.Run("Utxos", testUtxosAll)
}

func TestCount(t *testing.T) {
	t.Run("Blocks", testBlocksCount)
	t.Run("MempoolTransactions", testMempoolTransactionsCount)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsCount)
	t.Run("Utxos", testUtxosCount)
}

func TestHooks(t *testing.T) {
	t.Run("Blocks", testBlocksHooks)
	t.Run("MempoolTransactions", testMempoolTransactionsHooks)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsHooks)
	t.Run("Utxos", testUtxosHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Blocks", testBlocksInsert)
	t.Run("Blocks", testBlocksInsertWhitelist)
	t.Run("MempoolTransactions", testMempoolTransactionsInsert)
	t.Run("MempoolTransactions", testMempoolTransactionsInsertWhitelist)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsInsert)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsInsertWhitelist)
	t.Run("Utxos", testUtxosInsert)
	t.Run("Utxos", testUtxosInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Blocks", testBlocksReload)
	t.Run("MempoolTransactions", testMempoolTransactionsReload)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsReload)
	t.Run("Utxos", testUtxosReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Blocks", testBlocksReloadAll)
	t.Run("MempoolTransactions", testMempoolTransactionsReloadAll)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsReloadAll)
	t.Run("Utxos", testUtxosReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Blocks", testBlocksSelect)
	t.Run("MempoolTransactions", testMempoolTransactionsSelect)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsSelect)
	t.Run("Utxos", testUtxosSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Blocks", testBlocksUpdate)
	t.Run("MempoolTransactions", testMempoolTransactionsUpdate)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsUpdate)
	t.Run("Utxos", testUtxosUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Blocks", testBlocksSliceUpdateAll)
	t.Run("MempoolTransactions", testMempoolTransactionsSliceUpdateAll)
	t.Run("TransactionAuditlogs", testTransactionAuditlogsSliceUpdateAll)
	t.Run("Utxos", testUtxosSliceUpdateAll)
}
