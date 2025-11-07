package settings

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/db"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
)

var logger = modules.Logger{File: "app.log"}

func GetAccessToken() (string, error) {
	return db.GetValue("access-token")
}

func SetAccessToken(value string) error {
	return db.SetValue("access-token", value)
}
