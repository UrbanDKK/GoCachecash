// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testLogicalCacheMappings(t *testing.T) {
	t.Parallel()

	query := LogicalCacheMappings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLogicalCacheMappingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLogicalCacheMappingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LogicalCacheMappings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLogicalCacheMappingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LogicalCacheMappingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLogicalCacheMappingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LogicalCacheMappingExists(ctx, tx, o.EscrowID, o.SlotIdx)
	if err != nil {
		t.Errorf("Unable to check if LogicalCacheMapping exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LogicalCacheMappingExists to return true, but got false.")
	}
}

func testLogicalCacheMappingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	logicalCacheMappingFound, err := FindLogicalCacheMapping(ctx, tx, o.EscrowID, o.SlotIdx)
	if err != nil {
		t.Error(err)
	}

	if logicalCacheMappingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLogicalCacheMappingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LogicalCacheMappings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLogicalCacheMappingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LogicalCacheMappings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLogicalCacheMappingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	logicalCacheMappingOne := &LogicalCacheMapping{}
	logicalCacheMappingTwo := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, logicalCacheMappingOne, logicalCacheMappingDBTypes, false, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}
	if err = randomize.Struct(seed, logicalCacheMappingTwo, logicalCacheMappingDBTypes, false, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = logicalCacheMappingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = logicalCacheMappingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LogicalCacheMappings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLogicalCacheMappingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	logicalCacheMappingOne := &LogicalCacheMapping{}
	logicalCacheMappingTwo := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, logicalCacheMappingOne, logicalCacheMappingDBTypes, false, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}
	if err = randomize.Struct(seed, logicalCacheMappingTwo, logicalCacheMappingDBTypes, false, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = logicalCacheMappingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = logicalCacheMappingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func logicalCacheMappingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func logicalCacheMappingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *LogicalCacheMapping) error {
	*o = LogicalCacheMapping{}
	return nil
}

func testLogicalCacheMappingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &LogicalCacheMapping{}
	o := &LogicalCacheMapping{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping object: %s", err)
	}

	AddLogicalCacheMappingHook(boil.BeforeInsertHook, logicalCacheMappingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingBeforeInsertHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.AfterInsertHook, logicalCacheMappingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingAfterInsertHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.AfterSelectHook, logicalCacheMappingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingAfterSelectHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.BeforeUpdateHook, logicalCacheMappingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingBeforeUpdateHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.AfterUpdateHook, logicalCacheMappingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingAfterUpdateHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.BeforeDeleteHook, logicalCacheMappingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingBeforeDeleteHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.AfterDeleteHook, logicalCacheMappingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingAfterDeleteHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.BeforeUpsertHook, logicalCacheMappingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingBeforeUpsertHooks = []LogicalCacheMappingHook{}

	AddLogicalCacheMappingHook(boil.AfterUpsertHook, logicalCacheMappingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	logicalCacheMappingAfterUpsertHooks = []LogicalCacheMappingHook{}
}

func testLogicalCacheMappingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLogicalCacheMappingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(logicalCacheMappingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLogicalCacheMappingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLogicalCacheMappingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LogicalCacheMappingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLogicalCacheMappingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LogicalCacheMappings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	logicalCacheMappingDBTypes = map[string]string{`EscrowID`: `VARBINARY(16)`, `SlotIdx`: `UNSIGNED INT(4)`, `DatumID`: `VARBINARY(16)`}
	_                          = bytes.MinRead
)

func testLogicalCacheMappingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(logicalCacheMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(logicalCacheMappingColumns) == len(logicalCacheMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLogicalCacheMappingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(logicalCacheMappingColumns) == len(logicalCacheMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LogicalCacheMapping{}
	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LogicalCacheMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, logicalCacheMappingDBTypes, true, logicalCacheMappingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LogicalCacheMapping struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(logicalCacheMappingColumns, logicalCacheMappingPrimaryKeyColumns) {
		fields = logicalCacheMappingColumns
	} else {
		fields = strmangle.SetComplement(
			logicalCacheMappingColumns,
			logicalCacheMappingPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := LogicalCacheMappingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
