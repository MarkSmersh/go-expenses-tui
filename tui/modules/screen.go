package modules

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	View() string
	Update(tea.Msg) Cmd
	Keys() help.KeyMap
	GetTextInputs() []*textinput.Model
	SetActive()
	SetUnactive()
}
