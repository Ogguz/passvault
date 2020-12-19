package db

import (
	bolt "go.etcd.io/bbolt"
)

// Tx represents a BoltDB transaction
type Tx struct {
	*bolt.Tx
}

// Vault retrieves a Vault from the database with the given name.
func (tx *Tx) Vault(name []byte) (*Vault, error) {
	p := &Vault{
		Tx:   tx,
		Name: name,
	}

	return p, p.Load()
}
