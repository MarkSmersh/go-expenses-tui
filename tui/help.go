package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMapTransactionScreen struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
}

var KeysTransactionScreen = KeyMapTransactionScreen{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q", "ctrl+c"),
		key.WithHelp("ctrl+q", "quit"),
	),
}

func (k KeyMapTransactionScreen) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Up, k.Down,
		},
		{
			k.Quit,
		},
	}
}

func (k KeyMapTransactionScreen) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.Quit,
	}
}
