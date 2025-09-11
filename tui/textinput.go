package tui

import "github.com/charmbracelet/bubbles/textinput"

func CreateTextInput(placeholder string, limit int) textinput.Model {
	t := DefaultInput()

	t.Placeholder = placeholder
	t.CharLimit = limit

	return t
}
