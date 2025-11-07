package keys

import "github.com/charmbracelet/bubbles/key"

type KeyMapCommon struct {
	Quit       key.Binding
	PrevScreen key.Binding
	NextScreen key.Binding
}

var Common = KeyMapCommon{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q", "ctrl+c"),
		key.WithHelp("ctrl q", "quit"),
	),
	PrevScreen: key.NewBinding(
		key.WithKeys("ctrl+left"),
		key.WithHelp("ctrl ←", "previous screen"),
	),
	NextScreen: key.NewBinding(
		key.WithKeys("ctrl+right"),
		key.WithHelp("ctrl →", "next screen"),
	),
}

func (k KeyMapCommon) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Quit, k.PrevScreen, k.NextScreen,
		},
	}
}

func (k KeyMapCommon) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit, k.PrevScreen, k.NextScreen,
	}
}
