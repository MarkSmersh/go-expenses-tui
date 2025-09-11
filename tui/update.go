package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}

	case time.Time:
		return m, tick
	}

	m.Amount, cmd = m.Amount.Update(msg)

	return m, cmd
}

func tick() tea.Msg {
	time.Sleep(time.Second)
	return time.Time{}
}
