package main

// A simple program that counts down from 5 and then exits.
// actually yeah

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	model := tui.CreateModel()
	model.InitTextInputsFromScreens()

	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
