package modules

import tea "github.com/charmbracelet/bubbletea"

type Focusable interface {
	// View() string
	Focus() tea.Cmd
	Blur()
	Click()
}

type FocusManager struct {
	focused  int
	elements []Focusable
}

func CreateFocusManager(elements ...Focusable) FocusManager {
	return FocusManager{
		elements: elements,
		focused:  0,
	}
}

func (f *FocusManager) Set(cursor int) {
	if cursor > len(f.elements)-1 {
		f.focused = 0
	} else if cursor < 0 {
		f.focused = len(f.elements) - 1
	} else {
		f.focused = cursor
	}

	for _, e := range f.elements {
		e.Blur()
	}

	focused := f.elements[f.focused]
	focused.Focus()
}

func (f *FocusManager) Focused() Focusable {
	return f.elements[f.focused]
}

func (f *FocusManager) FocusedIndex() int {
	return f.focused
}

func (f *FocusManager) Next() {
	f.Set(f.focused + 1)
}

func (f *FocusManager) Prev() {
	f.Set(f.focused - 1)
}

func (f *FocusManager) BlurAll() {
	for _, e := range f.elements {
		e.Blur()
	}
}
