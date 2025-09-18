package screens

import (
	"fmt"

	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TransactionScreen struct {
	focused   int
	elements  int
	keys      KeyMapTransactionScreen
	focus     modules.FocusManager
	confirmed bool

	amount textinput.Model

	transactionType  textinput.Model
	transactionTypes []string

	comment textinput.Model
	confirm modules.Button
}

func NewTransactionScreen() *TransactionScreen {
	s := TransactionScreen{}

	s.focused = 0
	s.elements = 3
	s.keys = KeysTransactionScreen

	s.amount = modules.CreateTextInput("Amount", 32)
	s.transactionType = modules.CreateTextInput("Type of expense", 64)
	s.comment = modules.CreateTextInput("Comment", 128)

	s.transactionTypes = []string{"Stokrotka", "Gamling", "Onlyfans", "Woman", "Lawyer", "Random", "Sybau"}

	s.confirm = modules.NewButton("Confirm", func() {
		s.confirmed = !s.confirmed
	})

	var amount, transactionType, comment, confirm modules.Focusable

	// TODO: add to FocusManager a method that will create a clickable textinput from tea's input
	clickableAmount := modules.NewClickableTextInput(&s.amount)
	clickabletransactionType := modules.NewClickableTextInput(&s.transactionType)
	clickableComment := modules.NewClickableTextInput(&s.comment)

	amount = &clickableAmount
	transactionType = &clickabletransactionType
	comment = &clickableComment
	confirm = &s.confirm

	s.focus = modules.CreateFocusManager(amount, transactionType, comment, confirm)

	return &s
}

func (s TransactionScreen) View() string {
	view := "Create a new expense:"

	view += fmt.Sprintf("Current: %d", s.focused)

	view += "\n\n" + s.amount.View()
	view += "\n\n" + s.transactionType.View()
	view += "\n\n" + s.comment.View()
	view += "\n\n" + s.confirm.View()

	if s.confirmed {
		view += "\n\n" + "Congrats! Screen is confirmed. Throughout Heaven and Earth, I alone am the tui developer one"
	}

	return view
}

func (s *TransactionScreen) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd = nil
	k := s.keys

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, k.Quit):
			return tea.Quit
		case key.Matches(msg, k.Down):
			s.focus.Next()
		case key.Matches(msg, k.Up):
			s.focus.Prev()
		case key.Matches(msg, k.Activate):
			s.focus.Focused().Click()
		}

		s.focus.BlurAll()
		s.focus.Focused().Focus()
	}

	return cmd
}

func (s TransactionScreen) Keys() help.KeyMap {
	return s.keys
}

func (s *TransactionScreen) GetTextInputs() []*textinput.Model {
	tis := []*textinput.Model{
		&s.amount,
		&s.transactionType,
		&s.comment,
	}

	return tis
}
