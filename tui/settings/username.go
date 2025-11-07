package settings

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/db"
)

func GetUsername() (string, error) {
	return db.GetValue("username")
}

func SetUsername(value string) error {
	return db.SetValue("username", value)
}
