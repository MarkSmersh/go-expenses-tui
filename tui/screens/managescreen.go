package screens

import (
	"fmt"
	"slices"
	"time"

	"github.com/MarkSmersh/go-expenses-tui.git/tui/api"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/keys"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/modules"
	"github.com/MarkSmersh/go-expenses-tui.git/tui/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ManageScreen struct {
	block                 modules.BlockManager
	transactionsBlockKeys keys.KeyMapTransactionsBlock
	calenderBlockKeys     keys.KeyMapCalendarBlock
	chartBlockKeys        keys.KeyMapChartBlock
	style                 lipgloss.Style

	// Block 1
	transactions     modules.List
	transactionCache []api.TransactionExtended

	// Block 2
	calendar modules.Calendar

	// Block 3
	chart modules.Chart

	message string
	sum     int
}

func NewManageScreen() *ManageScreen {
	screen := ManageScreen{
		transactionsBlockKeys: keys.TransactionBlock,
		calenderBlockKeys:     keys.CalenderBlock,
		chartBlockKeys:        keys.ChartBlock,
		style:                 styles.Screen,

		transactions: modules.NewList(modules.NewListItem("No transactions?", "pupupu...")).
			WithShowHelp(false).
			WithTitle(false),

		calendar: modules.NewCalendar(),

		chart: modules.NewChart(6, 10,
			modules.NewChartItem("ras", 300),
			modules.NewChartItem("dvas", 600),
			modules.NewChartItem("tris", 900),
		).WithBarHeight(3).WithShowTotal(true),
	}

	screen.block = modules.NewBlockManager(
		modules.NewFocusManager(
			&screen.transactions,
		),
		modules.NewFocusManager(
			&screen.calendar,
		),
		modules.NewFocusManager(
			&screen.chart,
		),
	)

	return &screen
}

func (s ManageScreen) View() string {
	messageStyle := lipgloss.NewStyle().
		Bold(true).
		Width(s.style.GetWidth()-2).
		Height(1).
		Align(lipgloss.Left, lipgloss.Center)

	listStyle := lipgloss.NewStyle().
		Width(s.style.GetWidth()/2 - 2).
		// Height(s.style.GetHeight()-2).
		// Align(lipgloss.Left, lipgloss.Top).
		// Background(lipgloss.Color("3")).
		Border(lipgloss.HiddenBorder())

	calenderStyle := lipgloss.NewStyle().
		Width(s.style.GetWidth()/2-2).
		Height(8).
		// Background(lipgloss.Color("2")).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.HiddenBorder())

	summaryStyle := lipgloss.NewStyle().
		Width(s.style.GetWidth()/2 - 2).
		Height(s.style.GetHeight() - 6 - 8).
		// Background(lipgloss.Color("4")).
		Border(lipgloss.HiddenBorder())

	view := styles.ScreenTitle.
		Width(s.style.GetWidth() - 2).
		MarginBottom(1).
		Render("MANAGE SCREEN")

	view += lipgloss.JoinHorizontal(
		lipgloss.Top,
		listStyle.Render(
			s.transactions.View(),
		),
		lipgloss.JoinVertical(
			lipgloss.Right,
			calenderStyle.Render(
				s.calendar.View(),
			),
			summaryStyle.Render(
				s.chart.View(),
			),
		),
	)

	if len(s.message) > 0 {
		view += messageStyle.
			// Background(lipgloss.Color("5")).
			Render(s.message)
	}

	return s.style.Render(view)
}

func (s *ManageScreen) Update(msg tea.Msg) modules.Cmd {
	cmd := modules.NewCmd()
	k := s.Keys()

	transactionModel := s.transactions.Model()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s.message = ""

		switch k.(type) {
		case keys.KeyMapTransactionsBlock:
			if s.transactions.Model().FilterState() != list.Filtering {
				switch {
				case key.Matches(msg, s.transactionsBlockKeys.Tab):
					s.block.Next()
				case key.Matches(msg, s.transactionsBlockKeys.Delete):
					s.deleteTransaction(
						s.transactions.Model().Cursor(),
					)
					return cmd
				case key.Matches(msg, s.transactionsBlockKeys.More):
					// an information of transaction pop up window
				}
			}

		case keys.KeyMapCalendarBlock:
			switch {
			case key.Matches(msg, s.calenderBlockKeys.Tab):
				s.block.Next()
			case key.Matches(msg, s.calenderBlockKeys.Range):
				s.calendar.SwitchMode()
			case key.Matches(msg, s.calenderBlockKeys.Up):
				s.calendar.CursorUp()
			case key.Matches(msg, s.calenderBlockKeys.Down):
				s.calendar.CursorDown()
			case key.Matches(msg, s.calenderBlockKeys.Left):
				s.calendar.CursorLeft()
			case key.Matches(msg, s.calenderBlockKeys.Right):
				s.calendar.CursorRight()
			case key.Matches(msg, s.calenderBlockKeys.Selection):
				s.calendar.SwitchSelectionMode()
			case key.Matches(msg, s.calenderBlockKeys.Reset):
				s.calendar.Reset()
				s.updateTransactions()
				s.message = "Reset!"
			case key.Matches(msg, s.calenderBlockKeys.Update):
				if s.calendar.IsSelectionMode() {
					s.updateTransactions()
					s.message = "Updated from the calendar!"
				} else {
					s.message = "Updates from the callendar are only accepted in selection mode (c)."
				}
			}

		case keys.KeyMapChartBlock:
			switch {
			case key.Matches(msg, s.chartBlockKeys.Tab):
				s.block.Next()
			case key.Matches(msg, s.chartBlockKeys.Mark):
				s.chart.SwitchMarkSelected()
				s.filterTransactions()
			case key.Matches(msg, s.chartBlockKeys.UnmarkAll):
				s.chart.UnmarkAll()
				s.filterTransactions()
			case key.Matches(msg, s.chartBlockKeys.Reset):
				s.chart.MarkAll()
				s.filterTransactions()
			}
		}
	case tea.WindowSizeMsg:
		s.style = s.style.
			Width(msg.Width - 2).
			Height(msg.Height - 2 - 2)

		transactionModel.SetSize(s.style.GetWidth()/2-2, s.style.GetHeight()-4)

		s.chart.SetSize(s.style.GetHeight()-6-8, s.style.GetWidth()/2-2)
	}

	cmd.AddTea(
		s.transactions.Update(msg),
		s.chart.Update(msg),
	)

	return cmd
}

func (s ManageScreen) Keys() help.KeyMap {
	switch s.block.ActiveIndex() {
	case 0:
		return s.transactionsBlockKeys
	case 1:
		return s.calenderBlockKeys
	case 2:
		return s.chartBlockKeys
	}

	return nil
}

func (s *ManageScreen) SetActive() {
	s.updateTransactions()
	s.block.Focus()
}

func (s ManageScreen) SetUnactive() {
	s.block.BlurAll()
}

// fun fact! if you have no transactions and if you try to delete
// an item without selecting a one, it will crash. However, I do not
// ensure that it will happen, but I believe so!
func (s *ManageScreen) updateTransactions() {
	fromUnix := 0
	toUnix := 0

	if s.calendar.IsSelectionMode() {
		start, end := s.calendar.GetSelected()

		fromUnix = int(end.Unix())
		toUnix = int(start.Unix())

		logger.Logf("fromUnix: %d", fromUnix)
		logger.Logf("toUnix: %d", toUnix)
	}

	transactions, err := api.GetTransactions(100, fromUnix, toUnix, 0)

	s.transactionCache = transactions

	if err != nil {
		return
	}

	items := []modules.ListItem{}

	for _, t := range transactions {
		items = append(items, modules.NewListItem(
			fmt.Sprintf("%.2f zł | %s", float64(t.Amount)/100, t.TypeName),
			fmt.Sprintf("'%s'| %s", t.Comment, time.Unix(int64(t.CreatedAt), 0).Local().String()),
		))
	}

	s.transactions.SetItems(items...)

	s.updateChart()
}

func (s *ManageScreen) deleteTransaction(cursor int) {
	listLen := len(s.transactions.Model().Items())

	if listLen <= 0 {
		return
	}

	if err := api.DeleteTransaction(
		s.transactionCache[cursor].ID,
	); err != nil {
		s.message = fmt.Sprintf("Unable to delete the transaction. %s", err.Error())
	} else {
		s.message = "Deleted!"
	}

	s.updateTransactions()

	if listLen <= cursor+1 {
		s.transactions.Model().Select(cursor - 1)
	} else {
		s.transactions.Model().Select(cursor)
	}
}

func (s *ManageScreen) updateChart() {
	typesToAmount := map[string]float64{}

	for _, t := range s.transactionCache {
		_, ok := typesToAmount[t.TypeName]

		if ok {
			typesToAmount[t.TypeName] += float64(t.Amount / 100)
		} else {
			typesToAmount[t.TypeName] = float64(t.Amount / 100)
		}
	}

	items := []modules.ChartItem{}

	for t, a := range typesToAmount {
		items = append(items, modules.NewChartItem(t, a))
	}

	s.chart.SetItems(items...)
}

func (s *ManageScreen) filterTransactions() {

	items := []modules.ListItem{}

	marked := s.chart.GetMarked()

	markedTypes := []string{}

	for _, i := range marked {
		markedTypes = append(markedTypes, i.Name())
	}

	for _, t := range s.transactionCache {
		if slices.Contains(markedTypes, t.TypeName) {
			items = append(items, modules.NewListItem(
				fmt.Sprintf("%.2f zł | %s", float64(t.Amount)/100, t.TypeName),
				fmt.Sprintf("'%s'| %s", t.Comment, time.Unix(int64(t.CreatedAt), 0).Local().String()),
			))
		}
	}

	s.transactions.SetItems(items...)
}
