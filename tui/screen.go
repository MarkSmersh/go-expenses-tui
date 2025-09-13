package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	View() string
	// Init() Screen
	Focus(int)
	Focused() int
	Update(*Model, tea.Msg) tea.Cmd
	Keys() help.KeyMap
	GetTextInputs() []*textinput.Model
}

func (m Model) GetActiveScreen() Screen {
	return m.Screens[m.Screen]
}
