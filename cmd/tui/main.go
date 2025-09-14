package main

// A simple program that counts down from 5 and then exits.

import (
	"github.com/MarkSmersh/go-expenses-tui.git/tui"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	// logfilePath := os.Getenv("BUBBLETEA_LOG")
	// if logfilePath != "" {
	// 	if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	model := tui.CreateModel()
	model.InitTextInputsFromScreens()

	// Initialize our program
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
