package modules

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Button struct {
	focused bool
	text    string
	onClick func()
}

func NewButton(text string, onClick func()) Button {
	return Button{
		focused: false,
		text:    text,
		onClick: onClick,
	}
}

func (b *Button) View() string {
	if b.focused {
		return lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("[ %s ]", b.text))
	} else {
		return fmt.Sprintf("  %s  ", b.text)
	}
}

func (b *Button) Focus() tea.Cmd {
	b.focused = true
	return nil
}

func (b *Button) Blur() {
	b.focused = false
}

func (b *Button) Click() {
	b.onClick()
}
