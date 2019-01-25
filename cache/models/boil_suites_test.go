// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("GorpMigrations", testGorpMigrations)
	t.Run("LogicalCacheMappings", testLogicalCacheMappings)
}

func TestDelete(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsDelete)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsExists)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsExists)
}

func TestFind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsFind)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsFind)
}

func TestBind(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsBind)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsBind)
}

func TestOne(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsOne)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsOne)
}

func TestAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsAll)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsAll)
}

func TestCount(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsCount)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsCount)
}

func TestHooks(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsHooks)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsInsert)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsInsertWhitelist)
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
	t.Run("GorpMigrations", testGorpMigrationsReload)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSelect)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
	t.Run("LogicalCacheMappings", testLogicalCacheMappingsSliceUpdateAll)
}
