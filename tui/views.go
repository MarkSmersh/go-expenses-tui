package tui

import "fmt"

func (m Model) ExpenseScreenView() string {
	title := "Add a new expense"
	help := m.Help.View(m.Keys[0])

	return fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s", title, m.Amount.View(), m.MccCode.View(), help)
}
