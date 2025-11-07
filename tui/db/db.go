package db

import (
	"os"

	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	badger "github.com/dgraph-io/badger/v4"
)

var logger = modules.Logger{File: "app.log"}

func NewDB() *badger.DB {
	// FIXME: Will it work with the Windows FS?
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger").WithLogger(nil))

	if err != nil {
		logger.Logf("%s", err.Error())
		os.Exit(1)
	}

	return db
}
