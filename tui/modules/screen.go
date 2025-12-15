package modules

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	View() string
	Update(tea.Msg) Cmd
	Keys() help.KeyMap
	SetActive()
	SetUnactive()
}
