package tui

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	if !m.IsExclisive {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.Keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.Keys.PrevScreen):
				m.Screens.PrevScreen()
				return m, tea.WindowSize()
			case key.Matches(msg, m.Keys.NextScreen):
				m.Screens.NextScreen()
				return m, tea.WindowSize()
			}
		}
	}

	screenCmd := m.GetActiveScreen().Update(msg)

	switch screenCmd.GetScreen() {
	case modules.CmdExclusiveOff:
		m.SetDefaultScreens()
	case modules.CmdAuthScreen:
		m.SetExclusiveScreens()
	}

	cmds = append(cmds, screenCmd.GetTea()...)

	return m, tea.Batch(cmds...)
}
