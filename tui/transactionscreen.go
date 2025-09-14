package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TransactionScreen struct {
	focused  int
	elements int
	keys     KeyMapTransactionScreen

	amount textinput.Model

	transactionType  textinput.Model
	transactionTypes []string

	comment textinput.Model
}

func NewTransactionScreen() *TransactionScreen {
	s := TransactionScreen{}

	s.focused = 0
	s.elements = 3
	s.keys = KeysTransactionScreen

	s.amount = CreateTextInput("Amount", 32)

	s.transactionType = CreateTextInput("Type of expense", 64)

	s.comment = CreateTextInput("Comment", 128)

	s.comment.SetValue("test")

	s.transactionTypes = []string{"Stokrotka", "Gamling", "Onlyfans", "Woman", "Lawyer", "Random", "Sybau"}

	return &s
}

func (s TransactionScreen) View() string {
	view := "Create a new expense:"

	view += fmt.Sprintf("Current: %d", s.focused)

	view += "\n\n" + s.amount.View()
	view += "\n\n" + s.transactionType.View()
	view += "\n\n" + s.comment.View()

	if s.focused == 3 {
		view += "\n\n[ Confirm ]"
	} else {
		view += "\n\n  Confirm  "
	}

	return view
}

func (s TransactionScreen) Focused() int {
	return s.focused
}

func (s *TransactionScreen) Focus(i int) {
	if i > s.elements {
		s.focused = 0
	} else if i < 0 {
		s.focused = s.elements
	} else {
		s.focused = i
	}
}

func (s *TransactionScreen) Update(m *Model, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd = nil
	k := s.keys

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, k.Quit):
			return tea.Quit
		case key.Matches(msg, k.Down):
			s.Focus(s.Focused() + 1)
		case key.Matches(msg, k.Up):
			s.Focus(s.Focused() - 1)
		}

		s.amount.Blur()
		s.transactionType.Blur()
		s.comment.Blur()

		switch s.focused {
		case 0:
			s.amount.Focus()
		case 1:
			s.transactionType.Focus()
		case 2:
			s.comment.Focus()
		}
	}

	return cmd
}

func (s TransactionScreen) Keys() help.KeyMap {
	return s.keys
}

func (s TransactionScreen) GetTextInputs() []*textinput.Model {
	tis := []*textinput.Model{
		&s.amount,
		&s.transactionType,
		&s.comment,
	}

	return tis
}
