package tui

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/keys"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/screens"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/settings"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Help             help.Model
	Screens          modules.ScreenManager
	ExclusiveScreens modules.ScreenManager
	Logger           modules.Logger
	TextInputs       []*textinput.Model
	Keys             keys.KeyMapCommon
	// Switch to an exclisive screen mode. In difference to default ones,
	// exclusive ones are managed manually. Usefull in cases, when you need
	// to show one static screen. E.g. the Auth Screen.
	IsExclisive bool
}

func CreateModel() Model {
	m := Model{
		Screens: modules.NewScreenManager(
			screens.NewTransactionScreen(),
			screens.NewSettingsScreen(),
		),
		ExclusiveScreens: modules.NewScreenManager(
			screens.NewAuthScreen(),
		),
		Logger:      modules.Logger{File: "log"},
		Help:        help.New(),
		TextInputs:  []*textinput.Model{},
		Keys:        keys.Common,
		IsExclisive: false,
	}

	server, err := settings.GetServer()

	if err != nil || len(server) <= 0 {
		m.SetExclusiveScreens()
	}

	return m
}
