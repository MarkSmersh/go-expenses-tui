package modules

import "github.com/charmbracelet/bubbles/textinput"

func CreateTextInput(placeholder string, limit int) textinput.Model {
	t := DefaultInput()

	t.Placeholder = placeholder
	t.CharLimit = limit

	return t
}

func DefaultInput() textinput.Model {
	t := textinput.New()
	t.Placeholder = "Placeholder"
	t.CharLimit = 128
	t.Width = 30

	return t
}
