package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) AddTextInput(t *textinput.Model) {
	m.textInputs = append(m.textInputs, t)
}

func (m *Model) UpdateTextInputs(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{}

	for _, t := range m.textInputs {
		model, cmd := t.Update(msg)

		cmds = append(cmds, cmd)
		t = &model
	}

	return tea.Batch(cmds...)
}

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
