package db

import (
	bolt "go.etcd.io/bbolt"
	"time"
)

const path = "/opt/vault/passvault.db"

// DB represents a Bolt-backed data store.
type DB struct {
	*bolt.DB
}

// Open initializes and opens the database.
func (db *DB) Open() error {
	var err error

	db.DB, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		return err
	}

	// Create buckets.
	err = db.Update(func(tx *Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("vault")); err != nil {
			return &Error{"pages bucket error", err}
		}

		return nil
	})

	if err != nil {
		db.Close()

		return err
	}

	return nil
}

// View executes a function in the context of a read-only transaction.
func (db *DB) View(fn func(*Tx) error) error {
	return db.DB.View(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

// Update executes a function in the context of a writable transaction.
func (db *DB) Update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}