package modules

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	textInput textinput.Model
}

func NewInput(placeholder string, limit int) Input {
	t := DefaultTextInput()

	t.Placeholder = placeholder
	t.CharLimit = limit

	i := Input{
		textInput: t,
	}

	return i
}

func (c *Input) Click() {}

func (c *Input) Focus() tea.Cmd {
	c.textInput.Focus()
	return nil
}

func (c *Input) Blur() {
	c.textInput.Blur()
}

func (c Input) View() string {
	return c.textInput.View()
}

func (c *Input) TextInput() *textinput.Model {
	return &c.textInput
}

func (c Input) WithSecret() Input {
	c.textInput.EchoMode = textinput.EchoPassword
	c.textInput.EchoCharacter = '*'
	return c
}

func (c Input) WithSuggestions() Input {
	c.textInput.ShowSuggestions = true
	return c
}

func DefaultTextInput() textinput.Model {
	t := textinput.New()
	t.Placeholder = "Placeholder"
	t.CharLimit = 128
	t.Width = 30

	return t
}
