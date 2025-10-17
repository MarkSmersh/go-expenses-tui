package db

import "github.com/dgraph-io/badger/v4"

func GetValue(key string) (string, error) {
	db := NewDB()

	var server []byte

	err := db.View(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte("server"))

		if err != nil {
			return err
		}

		server, err = value.ValueCopy(nil)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return string(server), nil
}

func SetValue(key string, value string) error {
	db := NewDB()

	return db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
}
