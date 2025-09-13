package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	for _, s := range m.Screens {
		for _, ti := range s.GetTextInputs() {
			m.AddTextInput(ti)
		}
	}

	return textinput.Blink
}
