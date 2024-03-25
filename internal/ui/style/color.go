package style

import (
	"regexp"

	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

func ColorizeTable(row, col int) lipgloss.Style {
	switch {
	case col == 0:
		return ColorizeTabKey()
	}
	return lipgloss.NewStyle()
}

func measureContentWidth(agg int, item string, index int) int {
	itemSize := len(item)
	if itemSize > agg {
		return itemSize
	}
	return agg
}

func FormatTableWithFn(keys, content []string) func(int, int) lipgloss.Style {
	contentWidth := lo.Reduce(content, measureContentWidth, 10)
	keysWidth := lo.Reduce(keys, measureContentWidth, 20)

	return func(row, col int) lipgloss.Style {
		switch {
		case col == 0:
			return ColorizeTabKey().
				Width(contentWidth)
		case col == 1:
			return lipgloss.NewStyle().
				Width(keysWidth)
		}
		return lipgloss.NewStyle()
	}
}

func ColorizeTabKey() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(lipgloss.NoColor{}).
		Foreground(GetColor(ThemeConfig.Tab.ActiveTabContentKeys)).
		Padding(0, 1).
		Bold(true)
}

func GetColor(c string) lipgloss.TerminalColor {
	colorPattern, _ := regexp.Compile(`^#[0-9a-fA-F]{6}$`)
	if colorPattern.MatchString(c) {
		return lipgloss.Color(c)
	} else {
		return lipgloss.NoColor{}
	}
}

// func colorizeInt(string) string {
// 	return ""
// }

// func colorizeIP(string) string {
// 	return ""
// }

// func colorizeResourceType(string) string {
// 	return ""
// }

// func colorizeStatus(string) string {
// 	return ""
// }

// func colorizePort(string) string {
// 	return ""
// }

// func colorizeBool(string) string {
// 	return ""
// }
