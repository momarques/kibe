package style

import "github.com/charmbracelet/lipgloss"

var ()

func InSyncColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.InSyncState.BG)
}

func NotSyncedColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.NotSyncedState.BG)
}

func StartingColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.StartingState.BG)
}

func PausedColor() lipgloss.Color {
	return lipgloss.Color(ThemeConfig.SyncBar.PausedState.BG)
}

func SyncBarStatusStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(12)
}
