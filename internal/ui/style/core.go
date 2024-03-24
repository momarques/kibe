package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ViewTitleStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1)
)

func CoreHeaderTitleStyle() lipgloss.Style {
	return ViewTitleStyle.
		Foreground(GetColor(ThemeConfig.MainHeader.Title.TXT)).
		Background(GetColor(ThemeConfig.MainHeader.Title.BG)).
		MarginTop(1).
		MarginLeft(2)
}

func CoreHeaderItemCountStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.MainHeader.ItemCount.TXT)).
		MarginTop(1).
		MarginLeft(2).
		MarginBottom(1)
}

func HelpStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.Help.ShortcutName.TXT)).
		Background(GetColor(ThemeConfig.Help.ShortcutName.BG)).
		Padding(0, 0, 1, 0)
}
