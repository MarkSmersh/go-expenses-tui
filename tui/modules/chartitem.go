package modules

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ChartItem struct {
	name  string
	value float64
	// ASCII 256 colors
	color  lipgloss.Color
	marked bool
}

func NewChartItem(name string, value float64) ChartItem {
	color := 0

	for _, n := range name {
		color += int(n)

		if color > 256 {
			color = int(n)
		}
	}

	return ChartItem{
		name:   name,
		value:  value,
		color:  lipgloss.Color(fmt.Sprintf("%d", color)),
		marked: true,
	}
}

func (c ChartItem) FilterValue() string {
	return c.name
}

func (c *ChartItem) Mark(mark bool) {
	c.marked = mark
}

func (c ChartItem) Marked() bool {
	return c.marked
}

func (c ChartItem) Name() string {
	return c.name
}

type ChartItemDelegate struct{}

func (d ChartItemDelegate) Height() int { return 1 }

func (d ChartItemDelegate) Spacing() int { return 1 }

func (d ChartItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d ChartItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(ChartItem)

	if !ok {
		return
	}

	itemStyle := lipgloss.NewStyle().Foreground(i.color)

	str := ""

	if !i.marked {
		itemStyle = itemStyle.Strikethrough(true).Italic(true)
	}

	str += itemStyle.Render(i.name)

	if m.Index() == index {
		str = "> " + str
	}

	w.Write([]byte(str))
}
