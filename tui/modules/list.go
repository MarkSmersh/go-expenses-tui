package modules

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type List struct {
	list  list.Model
	items []ListItem
	focus bool
	color lipgloss.Color
}

func NewList(items ...ListItem) List {
	l := List{
		list: list.New(
			// why I can't unpack items while creating a slice?
			intoTeaItemList(items),
			list.NewDefaultDelegate(),
			0,
			0,
		),
		items: items,
		color: lipgloss.Color("5"),
	}

	return l
}

func (l List) Click() {}

func (l *List) Focus() tea.Cmd {
	l.focus = true
	return nil
}

func (l *List) Blur() {
	l.focus = false
}

func (l List) View() string {
	listStyle := lipgloss.NewStyle()

	if l.focus {
		listStyle = listStyle.Foreground(l.color)
	} else {
		listStyle = listStyle.Foreground(lipgloss.Color("2"))
	}

	return listStyle.Render(l.list.View())
}

func (l *List) Model() *list.Model {
	return &l.list
}

func (l List) WithTitleValue(title string) List {
	l.list.Title = title
	return l
}

func (l List) WithShowHelp(show bool) List {
	l.list.SetShowHelp(show)
	return l
}

// by default turn on
func (l List) WithTitle(show bool) List {
	l.list.SetShowTitle(show)
	return l
}

// shows metadata like items count of the list
func (l List) WithStatusBar(show bool) List {
	l.list.SetShowStatusBar(show)
	return l
}

func (l *List) AddItem(title string, description string) tea.Cmd {
	listItem := NewListItem(title, description)
	l.items = append(l.items, listItem)
	return l.list.InsertItem(-1, listItem)
}

func (l *List) SetItems(items ...ListItem) tea.Cmd {
	l.items = items

	return l.list.SetItems(
		intoTeaItemList(items),
	)
}

func (l *List) Update(msg tea.Msg) tea.Cmd {
	if !l.focus {
		return nil
	}

	var cmd tea.Cmd

	l.list, cmd = l.list.Update(msg)

	return cmd
}

func intoTeaItemList(items []ListItem) []list.Item {
	itemsList := []list.Item{}

	for _, i := range items {
		itemsList = append(itemsList, i)
	}

	return itemsList
}
