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
	for i, s := range m.Screens {
		m.Logger.Logf("SCREEN %d", i+1)

		for j, ti := range s.GetTextInputs() {
			m.Logger.Logf("TEXT INPUT %d", (i+1)*j+1)

			m.AddTextInput(ti)
		}
	}
}

func (m *Model) UpdateTextInputs(msg tea.Msg) tea.Cmd {
	cmds := []tea.Cmd{}

	m.Logger.Logf("inputs: %d", len(m.TextInputs))

	for _, t := range m.TextInputs {
		model, cmd := t.Update(msg)

		m.Logger.Logf("textinput %s, value %s, focus %t", model.Placeholder, model.Value(), model.Focused())

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
