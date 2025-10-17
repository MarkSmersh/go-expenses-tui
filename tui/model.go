package tui

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/screens"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Help       help.Model
	Screen     int
	Screens    []modules.Screen
	Logger     modules.Logger
	TextInputs []*textinput.Model
}

func CreateModel() Model {
	m := Model{
		Screen: 0,
		Screens: []modules.Screen{
			screens.NewTransactionScreen(),
		},
		Logger:     modules.Logger{File: "log"},
		Help:       help.New(),
		TextInputs: []*textinput.Model{},
	}

	return m
}
