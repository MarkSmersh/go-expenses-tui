package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func NewDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return db
}
