package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMapExensionScreen struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
}

var KeysExpensionScreen = KeyMapExensionScreen{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (k KeyMapExensionScreen) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Up, k.Down,
		},
		{
			k.Quit,
		},
	}
}

func (k KeyMapExensionScreen) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.Quit,
	}
}
