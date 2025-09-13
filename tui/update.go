package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	screenCmd := m.GetActiveScreen().Update(&m, msg)
	textInputsCmd := m.UpdateTextInputs(msg)

	cmds = append(cmds, screenCmd, textInputsCmd)

	return m, tea.Batch(cmds...)
}
