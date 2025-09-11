package tui

import "fmt"

func (m Model) View() string {
	view := m.Screens[m.Screen].View()
	help := m.Help.View(m.Keys[m.Screen])

	return fmt.Sprintf("%s\n\n%s", view, help)
}
