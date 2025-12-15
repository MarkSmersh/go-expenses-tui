package modules

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CalendarMode int

const (
	Day CalendarMode = iota
	Month
	Year
)

type Calendar struct {
	mode          CalendarMode
	selection     bool
	cursor, start time.Time
}

func NewCalendar() Calendar {
	t := time.Now()

	return Calendar{
		mode:      Day,
		cursor:    t,
		start:     t,
		selection: false,
	}
}

var logger = Logger{File: "app.log"}

func (c *Calendar) Click() {}

func (c *Calendar) Focus() tea.Cmd {
	return nil
}

func (c *Calendar) Blur() {
}

// FIXME: If the 2nd day of the month is Monday - the 1st day will be shifted
// and won't appear visually
func (c Calendar) View() string {
	calenderStyle := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center)

	view := ""

	y, m, d := c.cursor.Date()

	firstDayOfMonth := c.getDate(y, m, 0+1)
	dayCursor := 2 - int(firstDayOfMonth.Weekday())

	weeks := []string{}

	// 5 weeks multiplied by 7 days
	for range 5 {
		week := ""

		for range 7 {
			style := lipgloss.NewStyle()
			date := c.getDate(y, m, dayCursor)

			if date.Equal(c.GetCurrentDay()) {
				style = style.Bold(true)
			}

			if c.isDateSelected(date) {
				style = style.
					Foreground(lipgloss.Color("16")).
					Background(lipgloss.Color("15"))
			}

			dayString := strconv.Itoa(date.Day())

			// is selected day
			if date.Equal(c.getDate(y, m, d)) {
				week += style.Render(
					fmt.Sprintf("%4s",
						"["+fmt.Sprintf("%2s", dayString)+"]",
					),
				)
			} else {
				week += style.Render(
					fmt.Sprintf("%4s",
						" "+fmt.Sprintf("%2s", dayString)+" ",
					),
				)
			}

			dayCursor++
		}

		weeks = append(weeks, week)
	}

	selection := ""

	if c.selection {
		selection += fmt.Sprintf(
			"Selected: %d days",
			c.daysInBetween()+1,
		)
	}

	view += lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			fmt.Sprintf("%9s %d         %5s", m.String(), y, c.ModeString()),
		),
		lipgloss.NewStyle().Bold(true).Render(
			fmt.Sprintf(
				"%4s%4s%4s%4s%4s%4s%4s",
				"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun",
			),
		),
		strings.Join(weeks, "\n"),
		selection,
	)

	return calenderStyle.Render(view)
}

func (c *Calendar) SetSelectionMode(set bool) {
	if set {
		c.start = c.cursor
	}

	c.selection = set
}

func (c *Calendar) SwitchSelectionMode() {
	if c.selection {
		c.SetSelectionMode(false)
	} else {
		c.SetSelectionMode(true)
	}
}

func (c Calendar) IsSelectionMode() bool {
	return c.selection
}

func (c Calendar) getDate(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func (c Calendar) shiftDays(date time.Time, day int) time.Time {
	return c.getDate(date.Year(), date.Month(), date.Day()+day)
}

// days between start and cursor
func (c Calendar) daysInBetween() int {
	daySeconds := 60 * 60 * 24
	return int((c.start.Unix() - c.cursor.Unix()) / int64(daySeconds))
}

func (c Calendar) GetCurrentDay() time.Time {
	now := time.Now()
	return c.getDate(now.Year(), now.Month(), now.Day())
}

// returns the start and the cursor date in local time
func (c Calendar) GetSelected() (time.Time, time.Time) {
	y, m, d := c.start.Date()
	start := time.Date(y, m, d+1, 0, 0, 0-1, 0, time.Local)

	y, m, d = c.cursor.Date()
	cursor := time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	return start, cursor
}

func (c *Calendar) MoveCursor(days int) {
	c.cursor = c.shiftDays(c.cursor, days)
}

func (c Calendar) ModeString() string {
	switch c.mode {
	case Day:
		return "Day"
	case Month:
		return "Month"
	case Year:
		return "Year"
	}

	return ""
}

func (c *Calendar) CursorUp() {
	switch c.mode {
	case Day:
		c.MoveCursor(-7)
	case Month:
		c.MoveCursor(-30)
	case Year:
		c.MoveCursor(-365)
	}
}

func (c *Calendar) CursorDown() {
	switch c.mode {
	case Day:
		c.MoveCursor(+7)
	case Month:
		c.MoveCursor(+30)
	case Year:
		c.MoveCursor(+365)
	}
}

func (c *Calendar) CursorLeft() {
	c.MoveCursor(-1)
}

func (c *Calendar) CursorRight() {
	c.MoveCursor(+1)
}

// BLAZINGLY FAST
func (c *Calendar) SwitchMode() {
	switch c.mode {
	case Day:
		c.mode = Month
	case Month:
		c.mode = Year
	case Year:
		c.mode = Day
	}
}

func (c *Calendar) Reset() {
	c.cursor = c.GetCurrentDay()
	c.SetSelectionMode(false)
	c.mode = Day
}

func (c Calendar) isDateSelected(date time.Time) bool {
	if !c.selection {
		return false
	}

	start, end := c.GetSelected()

	if date.Unix() < start.Unix() && date.Unix() >= end.Unix() {
		return true
	} else {
		return false
	}
}
