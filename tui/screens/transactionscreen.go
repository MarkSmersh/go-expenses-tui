package screens

import (
	"math"
	"strconv"
	"strings"

	"github.com/MarkSmersh/go-expenses-tui.git/tui/api"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/keys"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var logger = modules.Logger{File: "app.log"}

type TransactionScreen struct {
	keys  keys.KeyMapTransactionScreen
	focus modules.FocusManager
	style lipgloss.Style

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
		transactionType: modules.NewInput("Type of expense (type a letter)", 64).WithSuggestions(),
		comment:         modules.NewInput("Comment", 128),

		transactionTypes: []api.TransactionType{},

		style: styles.Screen,
	}

	s.create = modules.NewButton("Create", s.createTransaction)

	s.updateTransactionTypesAndSuggestions()

	s.focus = modules.NewFocusManager(&s.amount, &s.transactionType, &s.comment, &s.create)

	return &s
}

func (s TransactionScreen) View() string {
	// heigth and width withouth margins and borders
	h := s.style.GetHeight() - 2
	w := s.style.GetWidth() - 2

	titleStyle := styles.ScreenTitle.Width(w).Height(1)

	inputStyle := lipgloss.NewStyle().Width(50).Height(2)

	transactionType := s.transactionType.TextInput().Value()
	suggestions := s.transactionType.TextInput().MatchedSuggestions()

	suggestionsView := []string{}

	blockStyle := lipgloss.NewStyle().
		Width(w).
		Height(h-2).
		Border(lipgloss.HiddenBorder()).
		// Background(lipgloss.Color("4")).
		Align(lipgloss.Center, lipgloss.Center)

	centerStyle := lipgloss.NewStyle().
		Width(50).
		Align(lipgloss.Center, lipgloss.Center)

	suggestionsStyle := lipgloss.NewStyle().
		Width(50)

	if len(transactionType) > 0 {
		if len(suggestions) > 0 {
			if transactionType != suggestions[0] || len(suggestions) > 1 {
				for _, s := range suggestions {
					suggestionsView = append(suggestionsView, s)
				}
			}
		} else {
			suggestionsView = append(suggestionsView, "No available suggestions")
		}
	}

	logger.Logf("SUGGESTIONS VIEW: %s", suggestionsView)

	return s.style.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		titleStyle.Render("TRANSACTION SCREEN"),
		blockStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				centerStyle.Render("Create a new expense:"),
				"\n",
				inputStyle.Render(s.amount.View()),
				inputStyle.Render(s.transactionType.View()),
				suggestionsStyle.Render(strings.Join(suggestionsView, "\n")),
				inputStyle.Render(s.comment.View()),
				centerStyle.Render(s.create.View()),
			),
		),
		s.message,
	))
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
	case tea.WindowSizeMsg:
		s.style = s.style.Height(msg.Height - 2 - 2).Width(msg.Width - 2)
	}

	cmd.AddTea(
		s.amount.Update(msg),
		s.transactionType.Update(msg),
		s.comment.Update(msg),
	)

	return cmd
}

func (s TransactionScreen) Keys() help.KeyMap {
	return s.keys
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

	transactionTypeIndex := -1

	for _, tt := range s.transactionTypes {
		if tt.Name == s.transactionType.TextInput().Value() {
			transactionTypeIndex = tt.ID
			break
		}
	}

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

	transactionTypeNames := []string{}

	for _, tt := range transactionTypes {
		transactionTypeNames = append(transactionTypeNames, tt.Name)
	}

	s.transactionTypes = transactionTypes
	s.transactionType.TextInput().SetSuggestions(transactionTypeNames)
}
