package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

func OKStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(false).
		Foreground(theme.GetColor(theme.Selected.StatusMessage.OKStatus.TXT))
}

func NOKStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.GetColor(theme.Selected.StatusMessage.NOKStatus.TXT))
}

func WarnStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.StatusMessage.WarnStatus.TXT))
}
