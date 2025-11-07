package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) AddTextInput(t ...*textinput.Model) {
	m.TextInputs = append(m.TextInputs, t...)
}

// FIXME: Needs to be rewrited. Too ugly to exist.
func (m *Model) InitTextInputsFromScreens() {
	for _, s := range m.Screens.GetScreens() {
		m.AddTextInput(s.GetTextInputs()...)
	}

	for _, s := range m.ExclusiveScreens.GetScreens() {
		m.AddTextInput(s.GetTextInputs()...)
	}
}

func (m *Model) UpdateTextInputs(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{}

	for _, t := range m.TextInputs {
		model, cmd := t.Update(msg)

		cmds = append(cmds, cmd)
		*t = model
	}

	return tea.Batch(cmds...)
}
