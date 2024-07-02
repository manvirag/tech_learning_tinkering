package main

import (
	"fmt"
)

// LegacyDatabase interface
type LegacyDatabase interface {
	SaveRecord(key, value string)
	LoadRecord(key string) string
}

// Standard Database interface
type Database interface {
	Insert(key, value string)
	Select(key string) string
}

// Adapter to convert LegacyDatabase to Database
type DatabaseAdapter struct {
	legacyDatabase LegacyDatabase
}

func NewDatabaseAdapter(legacyDatabase LegacyDatabase) *DatabaseAdapter {
	return &DatabaseAdapter{legacyDatabase: legacyDatabase}
}

func (d *DatabaseAdapter) Insert(key, value string) {
	d.legacyDatabase.SaveRecord(key, value)
}

func (d *DatabaseAdapter) Select(key string) string {
	return d.legacyDatabase.LoadRecord(key)
}

// LegacyDatabase implementation
type LegacyDatabaseImpl struct{}

func (ld *LegacyDatabaseImpl) SaveRecord(key, value string) {
	fmt.Printf("LegacyDatabaseImpl: Saving record with key %s and value %s\n", key, value)
}

func (ld *LegacyDatabaseImpl) LoadRecord(key string) string {
	return fmt.Sprintf("Loaded value for key %s", key)
}

func main() {
	legacyDatabase := &LegacyDatabaseImpl{}
	databaseAdapter := NewDatabaseAdapter(legacyDatabase)

	databaseAdapter.Insert("foo", "bar")
	value := databaseAdapter.Select("foo")
	fmt.Println(value)
}
