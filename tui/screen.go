package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Screen interface {
	View() string
	Init() Screen
	Focus(int)
	Focused() int
}

type TransactionScreen struct {
	focused  int
	elements int

	amount textinput.Model

	transactionType  textinput.Model
	transactionTypes []string

	comment textinput.Model
}

func (s TransactionScreen) Init() Screen {
	s.focused = 0
	s.elements = 3

	s.amount = CreateTextInput("Amount", 32)

	s.transactionType = CreateTextInput("Type of expense", 64)

	s.transactionTypes = []string{"Stokrotka", "Gamling", "Onlyfans", "Woman", "Lawyer", "Random", "Sybau"}

	return s
}

func (s TransactionScreen) View() string {
	view := "Create a new expense:"

	switch s.focused {
	case 0:
		s.amount.Focus()
	case 1:
		s.amount.Focus()
	case 2:
		s.amount.Focus()
	}

	view += "\n\n" + s.amount.View()
	view += "\n\n" + s.transactionType.View()
	view += "\n\n" + s.comment.View()

	if s.focused == 3 {
		view += "[Confirm]"
	} else {
		view += "Confirm"
	}

	return view
}

func (s TransactionScreen) Focused() int {
	return s.focused
}

func (s TransactionScreen) Focus(i int) {
	if i > s.elements {
		s.focused = 0
	} else {
		s.focused = i
	}
}
