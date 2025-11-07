package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapAuthScreen struct {
	Quit     key.Binding
	Up       key.Binding
	Down     key.Binding
	Activate key.Binding
}

var AuthScreen = KeyMapAuthScreen{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q"),
		key.WithHelp("ctrl q", "quit"),
	),
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

func (k KeyMapAuthScreen) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Quit, k.Up, k.Down, k.Activate,
		},
	}
}

func (k KeyMapAuthScreen) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit, k.Up, k.Down, k.Activate,
	}
}
