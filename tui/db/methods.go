package db

import (
	"github.com/dgraph-io/badger/v4"
)

func GetValue(key string) (string, error) {
	db := NewDB()

	defer db.Close()

	var server []byte

	err := db.View(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte(key))

		if err != nil {
			logger.Logf("Unable to get an %s", key)
			return err
		}

		server, err = value.ValueCopy(server)

		if err != nil {
			logger.Logf("Unable to get an %s", key)
			return err
		}

		return nil
	})

	if err != nil {
		logger.Logf("Unable to get an %s", key)
		return "", err
	}

	return string(server), nil
}

func SetValue(key string, value string) error {
	db := NewDB()

	defer db.Close()

	return db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
}

func Reset() error {
	db := NewDB()

	defer db.Close()

	err := db.DropAll()

	if err != nil {
		logger.Logf("Error while resettings settings. %s", err.Error())
	}

	return err
}
