package tui

import "fmt"

func (m Model) View() string {
	view := m.GetActiveScreen().View()

	help := m.Help.View(m.GetActiveScreen().Keys())

	if !m.IsExclisive {
		help += "\n" + m.Help.View(m.Keys)
	}

	return fmt.Sprintf("%s\n%s", view, help)
}
