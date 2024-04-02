package theme

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/samber/lo"
)

func measureContentWidth(agg int, item string, index int) int {
	itemSize := len(item)
	if itemSize > agg {
		return itemSize
	}
	return agg
}

func FormatTableWithFn(keys, content []string) func(int, int) lipgloss.Style {
	contentWidth := lo.Reduce(content, measureContentWidth, 20) + 2
	keysWidth := lo.Reduce(keys, measureContentWidth, 16) + 2

	return func(row, col int) lipgloss.Style {
		switch {
		case col == 0:
			return ColorizeTabKey().
				Width(keysWidth)
		case col == 1:
			return lipgloss.NewStyle().
				Width(contentWidth)
		}
		return lipgloss.NewStyle()
	}
}

func FormatTable(keys, content []string) string {
	rows := lo.Map(keys,
		func(item string, index int) []string {
			return []string{item + ":", content[index]}
		})

	t := table.New()
	t.Rows(rows...)
	t.StyleFunc(FormatTableWithFn(keys, content))
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func FormatSubTable(keys, content []string) string {
	keys = ColorizeSlice(keys)
	concatenated := lo.Map(keys, func(item string, index int) string {
		return fmt.Sprintf("%s -> %s", item, content[index])
	})
	return strings.Join(concatenated, "\n")
}

func ColorizeTabKey() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(lipgloss.NoColor{}).
		Foreground(GetColor(Selected.Tab.ActiveTabContentKeys)).
		PaddingLeft(1).
		Bold(true)
}

func ColorizeSlice(s []string) []string {
	return lo.Map(s,
		func(item string, _ int) string {
			return ColorizeTabKey().
				PaddingLeft(0).
				Bold(false).
				Render(item)
		})
}

func GetColor(c string) lipgloss.TerminalColor {
	colorPattern, _ := regexp.Compile(`^#[0-9a-fA-F]{6}$`)
	if colorPattern.MatchString(c) {
		return lipgloss.Color(c)
	} else {
		return lipgloss.NoColor{}
	}
}

func FormatCommand(cmds []string) string {
	return lipgloss.NewStyle().
		Italic(true).
		Foreground(GetColor(Selected.Tab.ActiveTabContentKeys)).
		Render(strings.Join(cmds, " "))
}
