package cmd

import "github.com/Ogguz/passvault/db"

// TODO add multi-db support
func newDB() (*db.DB, error)  {
	db := &db.DB{}

	if err := db.Open(); err != nil {
		return nil, err
	}

	return db, nil

}