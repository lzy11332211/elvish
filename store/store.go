package store

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

var createTable = map[string]string{}

// DefaultDB returns the default database for storage.
func DefaultDB() (*sql.DB, error) {
	ddir, err := ensureDataDir()
	if err != nil {
		return nil, err
	}
	uri := "file:" + url.QueryEscape(ddir+"/db") +
		"?mode=rwc&cache=shared&vfs=unix-dotfile"
	return sql.Open("sqlite3", uri)
}

// NewStore creates a new Store with the default database.
func NewStore() (*Store, error) {
	db, err := DefaultDB()
	if err != nil {
		return nil, err
	}
	return NewStoreDB(db)
}

// NewStoreDB creates a new Store with a custom database. The database must be
// a SQLite database.
func NewStoreDB(db *sql.DB) (*Store, error) {
	st := &Store{db}

	for t, q := range createTable {
		_, err := db.Exec(q)
		if err != nil {
			return nil, fmt.Errorf("failed to create table %s: %v", t, q)
		}
	}

	return st, nil
}
