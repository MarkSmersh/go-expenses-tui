package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapTransactionsBlock struct {
	Tab    key.Binding
	Up     key.Binding
	Down   key.Binding
	Delete key.Binding
	More   key.Binding
}

var TransactionBlock = KeyMapTransactionsBlock{
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("↹", "switch block"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "next"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "prev"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	More: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↩", "show more"),
	),
}

func (k KeyMapTransactionsBlock) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Tab, k.Down, k.Up, k.Delete, k.More,
		},
	}
}

func (k KeyMapTransactionsBlock) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Tab, k.Down, k.Up, k.Delete, k.More,
	}
}
