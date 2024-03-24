package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ViewTitleStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1)

	CoreHeaderTitleStyle = ViewTitleStyle.
				Copy().
				Foreground(GetColor(defaultThemeConfig.MainHeader.Title.TXT)).
				Background(GetColor(defaultThemeConfig.MainHeader.Title.BG)).
				MarginTop(1).
				MarginLeft(2)

	CoreHeaderItemCountStyle = lipgloss.NewStyle().
					Foreground(GetColor(defaultThemeConfig.MainHeader.ItemCount.TXT)).
					MarginTop(1).
					MarginLeft(2).
					MarginBottom(1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(GetColor(defaultThemeConfig.Help.ShortcutName.TXT)).
			Background(GetColor(defaultThemeConfig.Help.ShortcutName.BG)).
			Padding(0, 0, 1, 0)
)
