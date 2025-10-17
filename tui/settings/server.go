package settings

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/db"
)

func GetServer() (string, error) {
	return db.GetValue("server")
}

func SetServer(value string) error {
	return db.SetValue("server", value)
}
