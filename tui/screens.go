package tui

import "github.com/MarkSmersh/go-expenses-tui.git/tui/modules"

func (m Model) GetActiveScreen() modules.Screen {
	return m.Screens[m.Screen]
}
