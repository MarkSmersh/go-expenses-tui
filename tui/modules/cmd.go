package modules

import tea "github.com/charmbracelet/bubbletea"

const (
	CmdExclusiveOn  = 1
	CmdExclusiveOff = 2
	CmdAuthScreen   = 3
)

type Cmd struct {
	teaCmd    tea.Cmd
	screenCmd int
}

func NewCmd() Cmd {
	return Cmd{
		teaCmd:    nil,
		screenCmd: 0,
	}
}

func (c *Cmd) SetTea(teaCmd tea.Cmd) {
	c.teaCmd = teaCmd
}

func (c *Cmd) SetScreen(screenCmd int) {
	c.screenCmd = screenCmd
}

func (c Cmd) WithTea(teaCmd tea.Cmd) Cmd {
	c.teaCmd = teaCmd
	return c
}

func (c Cmd) WithScreen(screenCmd int) Cmd {
	c.screenCmd = screenCmd
	return c
}

func (c Cmd) GetTea() tea.Cmd {
	return c.teaCmd
}

func (c Cmd) GetScreen() int {
	return c.screenCmd
}
