package modules

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Chart struct {
	list          list.Model
	items         []ChartItem
	focus         bool
	height, width int
	barHeight     int
	totalHeight   int
}

func NewChart(height int, width int, items ...ChartItem) Chart {
	barHeight := 2

	slices.SortFunc(items, func(a, b ChartItem) int {
		return cmp.Compare(a.value, b.value) * -1
	})

	list := list.New(
		itemListToTeaList(items...),
		ChartItemDelegate{},
		width,
		height-barHeight,
	)

	list.SetShowTitle(false)
	list.SetShowStatusBar(false)
	list.SetShowHelp(false)

	return Chart{
		list:        list,
		items:       items,
		focus:       false,
		width:       20,
		height:      20,
		barHeight:   2,
		totalHeight: 0,
	}
}

func (c Chart) WithBarHeight(height int) Chart {
	c.barHeight = height

	c.list.SetHeight(c.height - height - c.totalHeight)

	return c
}

func (c Chart) WithShowTotal(show bool) Chart {
	if show {
		c.totalHeight = 1
	} else {
		c.totalHeight = 0
	}

	c.list.SetHeight(c.height - c.barHeight - c.totalHeight)

	return c
}

func (c *Chart) SetSize(height int, width int) {
	c.height = height
	c.width = width

	c.list.SetSize(width, height-c.barHeight-c.totalHeight)
}

func (c *Chart) Update(msg tea.Msg) tea.Cmd {
	if !c.focus {
		return nil
	}

	var cmd tea.Cmd

	c.list, cmd = c.list.Update(msg)

	return cmd
}

func (c Chart) Click() {}

func (c *Chart) Focus() tea.Cmd {
	c.focus = true
	return nil
}

func (c *Chart) Blur() {
	c.focus = false
}

func (c Chart) View() string {
	chartBar := []string{}

	markedItems := c.GetMarked()

	if len(markedItems) > 0 {
		values := []float64{}

		for _, i := range markedItems {
			values = append(values, i.value)
		}

		minValue := slices.Min(values)

		var valuesSum float64 = 0

		for i, v := range values {
			ampl := v / minValue
			values[i] = ampl
			valuesSum += ampl
		}

		amplWidth := float64(c.width) / valuesSum

		chartBarPart := []string{}

		for i, v := range values {
			for range int(v * amplWidth) {
				chartBarPart = append(chartBarPart,
					lipgloss.NewStyle().
						Foreground(markedItems[i].color).
						Render("█"),
				)
			}
		}

		if len(chartBarPart) < c.width {
			slices.Reverse(chartBarPart)

			for range c.width - len(chartBarPart) {
				chartBarPart = append(chartBarPart,
					lipgloss.NewStyle().
						Foreground(markedItems[0].color).
						Render("█"),
				)
			}

			slices.Reverse(chartBarPart)
		}

		for range c.barHeight {
			chartBar = append(chartBar, c.formatWithWidth(strings.Join(chartBarPart, "")))
		}

	} else {
		chartBar = make([]string, c.barHeight)

		chartBar[int(len(chartBar)/2)] = "No items..."
	}

	totalView := ""

	if c.totalHeight > 0 {
		var total float64 = 0

		for _, i := range markedItems {
			total += i.value
		}

		totalView += fmt.Sprintf("Total: %.2f zł", total)
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		c.list.View(),
		strings.Join(chartBar, "\n"),
		totalView,
	)
}

func (c *Chart) MarkSelected(mark bool) {
	c.GetItem().Mark(mark)
	c.updateList()
}

func (c *Chart) GetItem() *ChartItem {
	return &c.items[c.list.Index()]
}

func (c *Chart) SwitchMarkSelected() {
	item := c.GetItem()

	if item.Marked() {
		item.Mark(false)
	} else {
		item.Mark(true)
	}
	c.updateList()
}

func (c *Chart) UnmarkAll() {
	for i := range c.items {
		c.items[i].Mark(false)
	}

	c.updateList()
}

func (c *Chart) MarkAll() {
	for i := range c.items {
		c.items[i].Mark(true)
	}

	c.updateList()
}

func (c *Chart) SetItems(items ...ChartItem) {
	slices.SortFunc(items, func(a, b ChartItem) int {
		return cmp.Compare(a.value, b.value) * -1
	})
	c.items = items
	c.updateList()
}

func (c *Chart) GetMarked() []ChartItem {
	items := []ChartItem{}

	for _, i := range c.items {
		if i.Marked() {
			items = append(items, i)
		}
	}

	return items
}

func (c Chart) formatWithWidth(str string) string {
	widthStr := fmt.Sprintf("%d", c.width)

	return fmt.Sprintf("%"+widthStr+"s", str)
}

func (c *Chart) updateList() {
	c.list.SetItems(itemListToTeaList(c.items...))
}

func itemListToTeaList(items ...ChartItem) []list.Item {
	listItem := []list.Item{}

	for _, i := range items {
		listItem = append(listItem, i)
	}

	return listItem
}
