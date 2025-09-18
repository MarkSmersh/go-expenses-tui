package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) AddTextInput(t ...*textinput.Model) {
	m.Logger.Logf("ADDED TEXTINPUT. CURRENT COUNT OF TEXTINPUTS IS %d", len(m.TextInputs))
	m.TextInputs = append(m.TextInputs, t...)
}

func (m *Model) InitTextInputsFromScreens() {
	for _, s := range m.Screens {
		for _, ti := range s.GetTextInputs() {
			m.AddTextInput(ti)
		}
	}
}

func (m *Model) UpdateTextInputs(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{}

	m.Logger.Logf("inputs: %d", len(m.TextInputs))

	for _, t := range m.TextInputs {
		model, cmd := t.Update(msg)

		cmds = append(cmds, cmd)
		*t = model
	}

	return tea.Batch(cmds...)
}
