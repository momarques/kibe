package style

import "github.com/charmbracelet/lipgloss"

var (
	SyncBarStatusStyle = ViewTitleStyle.Copy().
				PaddingRight(0).
				Width(10)

	InSyncColor   = lipgloss.Color(defaultThemeConfig.SyncBar.InSyncState.BG)
	UnsyncedColor = lipgloss.Color(defaultThemeConfig.SyncBar.UnsyncedState.BG)
	StartingColor = lipgloss.Color(defaultThemeConfig.SyncBar.StartingState.BG)
)
