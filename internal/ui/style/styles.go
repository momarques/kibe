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

	SyncBarStatusStyle = ViewTitleStyle.Copy().
				PaddingRight(0).
				Width(10)

	InSyncColor   = defaultThemeConfig.SyncBar.InSyncState.TXT
	UnsyncedColor = defaultThemeConfig.SyncBar.UnsyncedState.TXT
	StartingColor = defaultThemeConfig.SyncBar.StartingState.TXT

	HelpStyle = lipgloss.NewStyle().
			Foreground(GetColor(defaultThemeConfig.Help.ShortcutName.TXT)).
			Background(GetColor(defaultThemeConfig.Help.ShortcutName.BG)).
			Padding(0, 0, 1, 0)
)
