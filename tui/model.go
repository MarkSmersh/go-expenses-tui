package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	Help    help.Model
	Screen  int
	Screens []Screen

	textInputs []*textinput.Model
}

func CreateModel() Model {
	m := Model{
		Screen: 0,
		Screens: []Screen{
			NewTransactionScreen(),
		},
		Help: help.New(),
	}

	return m
}
