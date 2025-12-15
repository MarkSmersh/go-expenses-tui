package styles

import "github.com/charmbracelet/lipgloss"

var Screen = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0)

var ScreenTitle = lipgloss.NewStyle().
	// Border(lipgloss.NormalBorder(), false, false, true, false).
	AlignHorizontal(lipgloss.Center).
	AlignHorizontal(lipgloss.Center).
	Bold(true).
	Height(1)
