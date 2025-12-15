package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapCalendarBlock struct {
	Tab       key.Binding
	Up        key.Binding
	Down      key.Binding
	Left      key.Binding
	Right     key.Binding
	Range     key.Binding
	Selection key.Binding
	Reset     key.Binding
	Update    key.Binding
}

var CalenderBlock = KeyMapCalendarBlock{
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
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "prev day"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "next day"),
	),
	Range: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "switch mode (day/month/year)"),
	),
	Selection: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "selection"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset"),
	),
	Update: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↩", "update"),
	),
}

func (k KeyMapCalendarBlock) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Tab, k.Down, k.Up, k.Left, k.Right, k.Range, k.Selection, k.Reset, k.Update,
		},
	}
}

func (k KeyMapCalendarBlock) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Tab, k.Down, k.Up, k.Left, k.Right, k.Range, k.Selection, k.Reset, k.Update,
	}
}
