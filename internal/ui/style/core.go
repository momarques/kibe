package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

var (
	ViewTitleStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1)
)

func CoreHeaderTitleStyle() lipgloss.Style {
	return ViewTitleStyle.
		Foreground(theme.GetColor(theme.Selected.MainHeader.Title.TXT)).
		Background(theme.GetColor(theme.Selected.MainHeader.Title.BG)).
		MarginLeft(2)
}

func CoreHeaderItemCountStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.MainHeader.ItemCount.TXT)).
		MarginTop(1).
		MarginLeft(2).
		MarginBottom(1)
}

func HelpStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.Help.ShortcutName.TXT)).
		Background(theme.GetColor(theme.Selected.Help.ShortcutName.BG)).
		Padding(0, 0, 1, 0)
}
