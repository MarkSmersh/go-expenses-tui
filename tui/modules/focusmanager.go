package modules

import tea "github.com/charmbracelet/bubbletea"

type Focusable interface {
	Focus() tea.Cmd
	Blur()
	// TODO: the Click method is used once: for the button module. And so,
	// it can be changed with the method Activate for the button, so Click method
	// could be removed from the interface
	Click()
}

type FocusManager struct {
	focused  int
	elements []Focusable
}

func NewFocusManager(elements ...Focusable) FocusManager {
	return FocusManager{
		elements: elements,
		focused:  0,
	}
}

func (f *FocusManager) Set(cursor int) {
	if len(f.elements) <= 0 {
		return
	}

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
