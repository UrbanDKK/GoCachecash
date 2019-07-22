// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// TransactionAuditlog is an object representing the database table.
type TransactionAuditlog struct {
	Rowid  int    `boil:"rowid" json:"rowid" toml:"rowid" yaml:"rowid"`
	Raw    []byte `boil:"raw" json:"raw" toml:"raw" yaml:"raw"`
	Status string `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *transactionAuditlogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L transactionAuditlogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TransactionAuditlogColumns = struct {
	Rowid  string
	Raw    string
	Status string
}{
	Rowid:  "rowid",
	Raw:    "raw",
	Status: "status",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var TransactionAuditlogWhere = struct {
	Rowid  whereHelperint
	Raw    whereHelper__byte
	Status whereHelperstring
}{
	Rowid:  whereHelperint{field: `rowid`},
	Raw:    whereHelper__byte{field: `raw`},
	Status: whereHelperstring{field: `status`},
}

// TransactionAuditlogRels is where relationship names are stored.
var TransactionAuditlogRels = struct {
}{}

// transactionAuditlogR is where relationships are stored.
type transactionAuditlogR struct {
}

// NewStruct creates a new relationship struct
func (*transactionAuditlogR) NewStruct() *transactionAuditlogR {
	return &transactionAuditlogR{}
}

// transactionAuditlogL is where Load methods for each relationship are stored.
type transactionAuditlogL struct{}

var (
	transactionAuditlogColumns               = []string{"rowid", "raw", "status"}
	transactionAuditlogColumnsWithoutDefault = []string{"raw", "status"}
	transactionAuditlogColumnsWithDefault    = []string{"rowid"}
	transactionAuditlogPrimaryKeyColumns     = []string{"rowid"}
)

type (
	// TransactionAuditlogSlice is an alias for a slice of pointers to TransactionAuditlog.
	// This should generally be used opposed to []TransactionAuditlog.
	TransactionAuditlogSlice []*TransactionAuditlog
	// TransactionAuditlogHook is the signature for custom TransactionAuditlog hook methods
	TransactionAuditlogHook func(context.Context, boil.ContextExecutor, *TransactionAuditlog) error

	transactionAuditlogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	transactionAuditlogType                 = reflect.TypeOf(&TransactionAuditlog{})
	transactionAuditlogMapping              = queries.MakeStructMapping(transactionAuditlogType)
	transactionAuditlogPrimaryKeyMapping, _ = queries.BindMapping(transactionAuditlogType, transactionAuditlogMapping, transactionAuditlogPrimaryKeyColumns)
	transactionAuditlogInsertCacheMut       sync.RWMutex
	transactionAuditlogInsertCache          = make(map[string]insertCache)
	transactionAuditlogUpdateCacheMut       sync.RWMutex
	transactionAuditlogUpdateCache          = make(map[string]updateCache)
	transactionAuditlogUpsertCacheMut       sync.RWMutex
	transactionAuditlogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var transactionAuditlogBeforeInsertHooks []TransactionAuditlogHook
var transactionAuditlogBeforeUpdateHooks []TransactionAuditlogHook
var transactionAuditlogBeforeDeleteHooks []TransactionAuditlogHook
var transactionAuditlogBeforeUpsertHooks []TransactionAuditlogHook

var transactionAuditlogAfterInsertHooks []TransactionAuditlogHook
var transactionAuditlogAfterSelectHooks []TransactionAuditlogHook
var transactionAuditlogAfterUpdateHooks []TransactionAuditlogHook
var transactionAuditlogAfterDeleteHooks []TransactionAuditlogHook
var transactionAuditlogAfterUpsertHooks []TransactionAuditlogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TransactionAuditlog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TransactionAuditlog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TransactionAuditlog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TransactionAuditlog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TransactionAuditlog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TransactionAuditlog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TransactionAuditlog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TransactionAuditlog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TransactionAuditlog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transactionAuditlogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTransactionAuditlogHook registers your hook function for all future operations.
func AddTransactionAuditlogHook(hookPoint boil.HookPoint, transactionAuditlogHook TransactionAuditlogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		transactionAuditlogBeforeInsertHooks = append(transactionAuditlogBeforeInsertHooks, transactionAuditlogHook)
	case boil.BeforeUpdateHook:
		transactionAuditlogBeforeUpdateHooks = append(transactionAuditlogBeforeUpdateHooks, transactionAuditlogHook)
	case boil.BeforeDeleteHook:
		transactionAuditlogBeforeDeleteHooks = append(transactionAuditlogBeforeDeleteHooks, transactionAuditlogHook)
	case boil.BeforeUpsertHook:
		transactionAuditlogBeforeUpsertHooks = append(transactionAuditlogBeforeUpsertHooks, transactionAuditlogHook)
	case boil.AfterInsertHook:
		transactionAuditlogAfterInsertHooks = append(transactionAuditlogAfterInsertHooks, transactionAuditlogHook)
	case boil.AfterSelectHook:
		transactionAuditlogAfterSelectHooks = append(transactionAuditlogAfterSelectHooks, transactionAuditlogHook)
	case boil.AfterUpdateHook:
		transactionAuditlogAfterUpdateHooks = append(transactionAuditlogAfterUpdateHooks, transactionAuditlogHook)
	case boil.AfterDeleteHook:
		transactionAuditlogAfterDeleteHooks = append(transactionAuditlogAfterDeleteHooks, transactionAuditlogHook)
	case boil.AfterUpsertHook:
		transactionAuditlogAfterUpsertHooks = append(transactionAuditlogAfterUpsertHooks, transactionAuditlogHook)
	}
}

// One returns a single transactionAuditlog record from the query.
func (q transactionAuditlogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TransactionAuditlog, error) {
	o := &TransactionAuditlog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for transaction_auditlog")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TransactionAuditlog records from the query.
func (q transactionAuditlogQuery) All(ctx context.Context, exec boil.ContextExecutor) (TransactionAuditlogSlice, error) {
	var o []*TransactionAuditlog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TransactionAuditlog slice")
	}

	if len(transactionAuditlogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TransactionAuditlog records in the query.
func (q transactionAuditlogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count transaction_auditlog rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q transactionAuditlogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if transaction_auditlog exists")
	}

	return count > 0, nil
}

// TransactionAuditlogs retrieves all the records using an executor.
func TransactionAuditlogs(mods ...qm.QueryMod) transactionAuditlogQuery {
	mods = append(mods, qm.From("\"transaction_auditlog\""))
	return transactionAuditlogQuery{NewQuery(mods...)}
}

// FindTransactionAuditlog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTransactionAuditlog(ctx context.Context, exec boil.ContextExecutor, rowid int, selectCols ...string) (*TransactionAuditlog, error) {
	transactionAuditlogObj := &TransactionAuditlog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"transaction_auditlog\" where \"rowid\"=$1", sel,
	)

	q := queries.Raw(query, rowid)

	err := q.Bind(ctx, exec, transactionAuditlogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from transaction_auditlog")
	}

	return transactionAuditlogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TransactionAuditlog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction_auditlog provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionAuditlogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	transactionAuditlogInsertCacheMut.RLock()
	cache, cached := transactionAuditlogInsertCache[key]
	transactionAuditlogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			transactionAuditlogColumns,
			transactionAuditlogColumnsWithDefault,
			transactionAuditlogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(transactionAuditlogType, transactionAuditlogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(transactionAuditlogType, transactionAuditlogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"transaction_auditlog\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"transaction_auditlog\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into transaction_auditlog")
	}

	if !cached {
		transactionAuditlogInsertCacheMut.Lock()
		transactionAuditlogInsertCache[key] = cache
		transactionAuditlogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TransactionAuditlog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TransactionAuditlog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	transactionAuditlogUpdateCacheMut.RLock()
	cache, cached := transactionAuditlogUpdateCache[key]
	transactionAuditlogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			transactionAuditlogColumns,
			transactionAuditlogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update transaction_auditlog, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"transaction_auditlog\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, transactionAuditlogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(transactionAuditlogType, transactionAuditlogMapping, append(wl, transactionAuditlogPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update transaction_auditlog row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for transaction_auditlog")
	}

	if !cached {
		transactionAuditlogUpdateCacheMut.Lock()
		transactionAuditlogUpdateCache[key] = cache
		transactionAuditlogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q transactionAuditlogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for transaction_auditlog")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for transaction_auditlog")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TransactionAuditlogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionAuditlogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"transaction_auditlog\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, transactionAuditlogPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in transactionAuditlog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all transactionAuditlog")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TransactionAuditlog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction_auditlog provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionAuditlogColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	transactionAuditlogUpsertCacheMut.RLock()
	cache, cached := transactionAuditlogUpsertCache[key]
	transactionAuditlogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			transactionAuditlogColumns,
			transactionAuditlogColumnsWithDefault,
			transactionAuditlogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			transactionAuditlogColumns,
			transactionAuditlogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert transaction_auditlog, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(transactionAuditlogPrimaryKeyColumns))
			copy(conflict, transactionAuditlogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"transaction_auditlog\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(transactionAuditlogType, transactionAuditlogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(transactionAuditlogType, transactionAuditlogMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert transaction_auditlog")
	}

	if !cached {
		transactionAuditlogUpsertCacheMut.Lock()
		transactionAuditlogUpsertCache[key] = cache
		transactionAuditlogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TransactionAuditlog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TransactionAuditlog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransactionAuditlog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), transactionAuditlogPrimaryKeyMapping)
	sql := "DELETE FROM \"transaction_auditlog\" WHERE \"rowid\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from transaction_auditlog")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for transaction_auditlog")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q transactionAuditlogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no transactionAuditlogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transaction_auditlog")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction_auditlog")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TransactionAuditlogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransactionAuditlog slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(transactionAuditlogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionAuditlogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"transaction_auditlog\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionAuditlogPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transactionAuditlog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction_auditlog")
	}

	if len(transactionAuditlogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TransactionAuditlog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTransactionAuditlog(ctx, exec, o.Rowid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionAuditlogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TransactionAuditlogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionAuditlogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"transaction_auditlog\".* FROM \"transaction_auditlog\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionAuditlogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TransactionAuditlogSlice")
	}

	*o = slice

	return nil
}

// TransactionAuditlogExists checks if the TransactionAuditlog row exists.
func TransactionAuditlogExists(ctx context.Context, exec boil.ContextExecutor, rowid int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"transaction_auditlog\" where \"rowid\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, rowid)
	}

	row := exec.QueryRowContext(ctx, sql, rowid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if transaction_auditlog exists")
	}

	return exists, nil
}
