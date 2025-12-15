package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapChartBlock struct {
	Tab       key.Binding
	Up        key.Binding
	Down      key.Binding
	Mark      key.Binding
	UnmarkAll key.Binding
	Reset     key.Binding
}

var ChartBlock = KeyMapChartBlock{
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("↹", "switch block"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "next period"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "prev period"),
	),
	Mark: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↩", "mark"),
	),
	UnmarkAll: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "clear"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset"),
	),
}

func (k KeyMapChartBlock) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Tab, k.Down, k.Up, k.Mark, k.UnmarkAll, k.Reset,
		},
	}
}

func (k KeyMapChartBlock) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Tab, k.Down, k.Up, k.Mark, k.UnmarkAll, k.Reset,
	}
}
