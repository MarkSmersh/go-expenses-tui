package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Keys []help.KeyMap
	Help help.Model

	Screen int

	Screens []Screen

	Amount  textinput.Model
	MccCode textinput.Model
}

func CreateModel() Model {
	m := Model{
		Screen:  0,
		Amount:  CreateTextInput("Amount", 128),
		MccCode: CreateTextInput("MCC code (optional)", 4),
		Keys: []help.KeyMap{
			KeysExpensionScreen,
		},
		Screens: []Screen{
			ExpenseScreen{},
		},
		Help: help.New(),
	}

	return m
}

func DefaultInput() textinput.Model {
	t := textinput.New()
	t.Placeholder = "Placeholder"
	t.CharLimit = 128
	t.Width = 30

	return t
}
