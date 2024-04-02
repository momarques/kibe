package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

var ()

func InSyncColor() lipgloss.Color {
	return lipgloss.Color(theme.Selected.SyncBar.InSyncState.BG)
}

func NotSyncedColor() lipgloss.Color {
	return lipgloss.Color(theme.Selected.SyncBar.NotSyncedState.BG)
}

func StartingColor() lipgloss.Color {
	return lipgloss.Color(theme.Selected.SyncBar.StartingState.BG)
}

func PausedColor() lipgloss.Color {
	return lipgloss.Color(theme.Selected.SyncBar.PausedState.BG)
}

func SyncBarStatusStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(12)
}
