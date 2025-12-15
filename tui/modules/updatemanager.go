package modules

import tea "github.com/charmbracelet/bubbletea"

type Updatable interface {
	Update() tea.Cmd
}

type UpdateManager struct {
	elements []*Updatable
}
