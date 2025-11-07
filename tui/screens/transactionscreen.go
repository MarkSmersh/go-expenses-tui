package screens

import (
	"math"
	"strconv"

	"github.com/MarkSmersh/go-expenses-tui.git/tui/api"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/keys"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var logger = modules.Logger{File: "app.log"}

type TransactionScreen struct {
	keys  keys.KeyMapTransactionScreen
	focus modules.FocusManager

	amount          modules.Input
	transactionType modules.Input
	comment         modules.Input
	create          modules.Button

	transactionTypes []api.TransactionType
	message          string
}

func NewTransactionScreen() *TransactionScreen {
	s := TransactionScreen{
		keys:            keys.TransactionScreen,
		amount:          modules.NewInput("Amount (required)", 32),
		transactionType: modules.NewInput("Type of expense (write first letter to show suggestions)", 64).WithSuggestions(),
		comment:         modules.NewInput("Comment", 128),

		transactionTypes: []api.TransactionType{},
	}

	s.create = modules.NewButton("Create", s.createTransaction)

	s.updateTransactionTypesAndSuggestions()

	logger.Logf("SUGGESTIONS: %v", s.transactionType.TextInput().AvailableSuggestions())

	s.focus = modules.CreateFocusManager(&s.amount, &s.transactionType, &s.comment, &s.create)

	return &s
}

func (s TransactionScreen) View() string {
	view := "TRANSACTION SCREEN"

	view += "\n\nCreate a new expense:"

	view += "\n\n" + s.amount.View()
	view += "\n\n" + s.transactionType.View()

	transactionType := s.transactionType.TextInput().Value()
	suggestions := s.transactionType.TextInput().MatchedSuggestions()

	if len(transactionType) > 0 && len(suggestions) > 0 && transactionType != suggestions[0] {
		view += "\n"

		for _, s := range suggestions {
			view += "\n" + s
		}

		if len(suggestions) <= 0 {
			view += "No available suggestions"
		}
	}

	view += "\n\n" + s.comment.View()
	view += "\n\n" + s.create.View()

	if len(s.message) > 0 {
		view += " - " + s.message
	}

	return view
}

func (s *TransactionScreen) Update(msg tea.Msg) modules.Cmd {
	cmd := modules.NewCmd()
	k := s.keys

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s.message = ""

		switch {
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
		s.amount.TextInput(),
		s.transactionType.TextInput(),
		s.comment.TextInput(),
	}

	return tis
}

func (s *TransactionScreen) SetActive() {
	s.updateTransactionTypesAndSuggestions()
	s.focus.Focused().Focus()
}

func (s TransactionScreen) SetUnactive() {
	s.focus.BlurAll()
}

func (s *TransactionScreen) createTransaction() {
	amountFloat, err := strconv.ParseFloat(s.amount.TextInput().Value(), 64)

	if err != nil {
		s.message = "the amount field has non-number value"
		return
	}

	if len(s.transactionType.TextInput().Value()) < 1 {
		s.message = "the type of expense field cannot be empty"
		return
	}

	transactionTypeIndex := s.transactionType.TextInput().CurrentSuggestionIndex() + 1
	comment := s.comment.TextInput().Value()

	amount := int(math.Round(amountFloat * 100))

	err = api.CreateTransaction(amount, transactionTypeIndex, comment)

	if err != nil {
		s.message = err.Error()
		return
	}

	s.amount.TextInput().SetValue("")
	s.transactionType.TextInput().SetValue("")
	s.comment.TextInput().SetValue("")

	s.message = "transaction has been created succesfully!"
}

func (s *TransactionScreen) updateTransactionTypesAndSuggestions() {
	transactionTypes, err := api.GetTransactionTypes()

	if err != nil {
		logger.Logf("Trying to access transaction types from Transaction Screen. %s", err.Error())
		// os.Exit(1)
	}

	logger.Logf("TYPES: %v", transactionTypes)

	transactionTypeNames := []string{}

	for _, tt := range transactionTypes {
		transactionTypeNames = append(transactionTypeNames, tt.Name)
	}

	s.transactionType.TextInput().SetSuggestions(transactionTypeNames)
}
