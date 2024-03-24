package style

import "github.com/charmbracelet/lipgloss"

var ()

func InSyncColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.InSyncState.BG)
}

func UnsyncedColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.UnsyncedState.BG)
}

func StartingColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.StartingState.BG)
}

func SyncBarStatusStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(0).
		Width(10)
}
