package modules

import tea "github.com/charmbracelet/bubbletea"

const (
	CmdExclusiveOn  = 1
	CmdExclusiveOff = 2
	CmdAuthScreen   = 3
)

type Cmd struct {
	teaCmds   []tea.Cmd
	screenCmd int
}

func NewCmd() Cmd {
	return Cmd{
		teaCmds:   []tea.Cmd{},
		screenCmd: 0,
	}
}

func (c *Cmd) AddTea(teaCmd ...tea.Cmd) {
	c.teaCmds = append(c.teaCmds, teaCmd...)
}

func (c *Cmd) SetScreen(screenCmd int) {
	c.screenCmd = screenCmd
}

func (c Cmd) WithTea(teaCmd ...tea.Cmd) Cmd {
	c.teaCmds = teaCmd
	return c
}

func (c Cmd) WithScreen(screenCmd int) Cmd {
	c.screenCmd = screenCmd
	return c
}

func (c Cmd) GetTea() []tea.Cmd {
	return c.teaCmds
}

func (c Cmd) GetScreen() int {
	return c.screenCmd
}
