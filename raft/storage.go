// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
// Maintained by : Rukshan Perera (rukshan.perera@student.oulu.fi)
package raft

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"os"
	"sync"
)

// Storage is an interface implemented by stable storage providers.
type Storage interface {
	Set(key string, value []byte)

	Get(key string) ([]byte, bool)

	// HasData returns true iff any Sets were made on this Storage.
	HasData() bool
}

// SQLiteStorage is an implementation of Storage using SQLite.
type SQLiteStorage struct {
	mu     sync.Mutex
	dbPath string
}

// MapStorage is a simple in-memory implementation of Storage for testing.
type MapStorage struct {
	mu sync.Mutex
	m  map[string][]byte
}

func NewSQLiteStorage(dbPath string) *SQLiteStorage {
	_ = os.MkdirAll("./", 0755)
	_, _ = os.Create(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Error openning database: %v\n", err)
		return nil
	}
	// Ensure that the 'data' table exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS data (key TEXT PRIMARY KEY, value BLOB)")
	if err != nil {
		fmt.Printf("table creation error: %v\n", err)
		return nil
	}
	_ = db.Close()
	return &SQLiteStorage{dbPath: dbPath}
}

// Set stores the given key-value pair in the SQLite database.
func (ss *SQLiteStorage) Set(key string, value []byte) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	db, err := sql.Open("sqlite3", ss.dbPath)
	if err != nil {
		fmt.Printf("Error openning database: %v\n", err)
		return
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Error beginning transaction: %v\n", err)
		_ = db.Close()
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT OR REPLACE INTO data (key, value) VALUES (?, ?)", key, value)
	if err != nil {
		fmt.Printf("Error inserting data: %v\n", err)
		_ = db.Close()
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error committing transaction: %v\n", err)
	}
	_ = db.Close()
}

// Get retrieves the value associated with the given key from the SQLite database.
func (ss *SQLiteStorage) Get(key string) ([]byte, bool) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	db, err := sql.Open("sqlite3", ss.dbPath)
	if err != nil {
		fmt.Printf("Error openning database: %v\n", err)
		err := db.Close()
		if err != nil {
			return nil, false
		}
		return nil, false
	}
	var value []byte
	err = db.QueryRow("SELECT value FROM data WHERE key = ?", key).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			_ = db.Close()
			return nil, false
		}
		fmt.Printf("Error retrieving data: %v\n", err)
		_ = db.Close()
		return nil, false
	}
	_ = db.Close()
	return value, true
}

// HasData checks if there is any data stored in the SQLite database.
func (ss *SQLiteStorage) HasData() bool {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	db, err := sql.Open("sqlite3", ss.dbPath)
	if err != nil {
		fmt.Printf("Error openning database: %v\n", err)
		return false
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM data").Scan(&count)
	if err != nil {
		fmt.Printf("Error checking data: %v\n", err)
		_ = db.Close()
		return false
	}
	_ = db.Close()
	return count > 0
}

// Close closes the SQLite database connection.
func (ss *SQLiteStorage) Close() error {
	return nil
}

func NewMapStorage() *MapStorage {
	m := make(map[string][]byte)
	return &MapStorage{
		m: m,
	}
}

func (ms *MapStorage) Get(key string) ([]byte, bool) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	v, found := ms.m[key]
	return v, found
}

func (ms *MapStorage) Set(key string, value []byte) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.m[key] = value
}

func (ms *MapStorage) HasData() bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return len(ms.m) > 0
}
