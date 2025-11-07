package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapTransactionScreen struct {
	Up       key.Binding
	Down     key.Binding
	Activate key.Binding
}

var TransactionScreen = KeyMapTransactionScreen{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	Activate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "activate"),
	),
}

func (k KeyMapTransactionScreen) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Up, k.Down, k.Activate,
		},
	}
}

func (k KeyMapTransactionScreen) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.Activate,
	}
}
