package tui

import "github.com/MarkSmersh/go-expenses-tui.git/tui/modules"

func (m Model) GetActiveScreen() modules.Screen {
	if m.IsExclisive {
		return m.ExclusiveScreens.GetActiveScreen()
	} else {
		return m.Screens.GetActiveScreen()
	}
}

func (m *Model) SetExclusiveScreens() {
	m.GetActiveScreen().SetUnactive()
	m.IsExclisive = true
	m.ExclusiveScreens.SetActiveScreen(
		m.ExclusiveScreens.GetActiveScreenIndex(),
	)
}

func (m *Model) SetDefaultScreens() {
	m.GetActiveScreen().SetUnactive()
	m.IsExclisive = false
	m.Screens.SetActiveScreen(
		m.Screens.GetActiveScreenIndex(),
	)
}
