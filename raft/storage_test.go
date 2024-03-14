package raft

import (
	"testing"
)

func TestNewSQLiteStorage(t *testing.T) {
	dbPath := "./test.db"
	storage := NewSQLiteStorage(dbPath)
	if storage == nil {
		t.Errorf("Failed to create new SQLite storage")
	}
}

func TestSQLiteStorage_Set_Get(t *testing.T) {
	dbPath := "./test.db"
	storage := NewSQLiteStorage(dbPath)
	defer storage.Close()

	key := "testKey"
	value := []byte("testValue")

	// Test Set
	storage.Set(key, value)

	// Test Get
	retrievedValue, ok := storage.Get(key)
	if !ok || string(retrievedValue) != string(value) {
		t.Errorf("Failed to get the correct value. Expected %s, got %s", string(value), string(retrievedValue))
	}
}

func TestSQLiteStorage_HasData(t *testing.T) {
	dbPath := "./test.db"
	storage := NewSQLiteStorage(dbPath)
	defer storage.Close()

	key := "testKey"
	value := []byte("testValue")
	storage.Set(key, value)

	if !storage.HasData() {
		t.Errorf("HasData returned false, expected true")
	}
}

func TestSQLiteStorage_ReadWrite(t *testing.T) {
	dbPath := "./test.db"
	var storage = NewSQLiteStorage(dbPath)

	var key = "testKey"
	var value = []byte("testValue")
	storage.Set(key, value)

	var retrievedValue, ok = storage.Get(key)
	if !ok || string(retrievedValue) != string(value) {
		t.Errorf("Failed to get the correct value. Expected %s, got %s", string(value), string(retrievedValue))
	}

	key = "testKey2"
	value = []byte("testValue2")
	storage.Set(key, value)

	retrievedValue, ok = storage.Get(key)
	if !ok || string(retrievedValue) != string(value) {
		t.Errorf("Failed to get the correct value. Expected %s, got %s", string(value), string(retrievedValue))
	}
	storage.Close()
}
