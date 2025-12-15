package modules

type ListItem struct {
	title, description string
}

func NewListItem(title string, description string) ListItem {
	return ListItem{
		title:       title,
		description: description,
	}
}

func (i ListItem) Title() string       { return i.title }
func (i ListItem) Description() string { return i.description }
func (i ListItem) FilterValue() string { return i.title }
