package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

func InfoLevel() lipgloss.TerminalColor {
	return theme.GetColor(theme.Selected.StatusLog.InfoLevel.TXT)
}
func WarnLevel() lipgloss.TerminalColor {
	return theme.GetColor(theme.Selected.StatusLog.WarnLevel.TXT)
}
func ErrorLevel() lipgloss.TerminalColor {
	return theme.GetColor(theme.Selected.StatusLog.ErrorLevel.TXT)
}
func DebugLevel() lipgloss.TerminalColor {
	return theme.GetColor(theme.Selected.StatusLog.DebugLevel.TXT)
}

func StatusLogDuration() lipgloss.TerminalColor {
	return theme.GetColor(theme.Selected.StatusLog.Duration.TXT)
}

func OKStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(false).
		Foreground(theme.GetColor(theme.Selected.StatusLog.OKStatus.TXT))
}

func NOKStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.GetColor(theme.Selected.StatusLog.NOKStatus.TXT))
}

func WarnStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.StatusLog.WarnStatus.TXT))
}

func NoneStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.StatusLog.Duration.TXT))
}
