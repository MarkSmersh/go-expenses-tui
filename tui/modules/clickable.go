package modules

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ClickableTextInput struct {
	textInput *textinput.Model
}

func NewClickableTextInput(textInput *textinput.Model) ClickableTextInput {
	return ClickableTextInput{
		textInput: textInput,
	}
}

func (c *ClickableTextInput) Click() {}

func (c *ClickableTextInput) Focus() tea.Cmd {
	c.textInput.Focus()
	return nil
}

func (c *ClickableTextInput) Blur() {
	c.textInput.Blur()
}
