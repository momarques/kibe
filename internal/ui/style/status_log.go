package style

import "github.com/charmbracelet/lipgloss"

func InfoLevel() lipgloss.TerminalColor  { return GetColor(ThemeConfig.StatusLog.InfoLevel.TXT) }
func WarnLevel() lipgloss.TerminalColor  { return GetColor(ThemeConfig.StatusLog.WarnLevel.TXT) }
func ErrorLevel() lipgloss.TerminalColor { return GetColor(ThemeConfig.StatusLog.ErrorLevel.TXT) }
func DebugLevel() lipgloss.TerminalColor { return GetColor(ThemeConfig.StatusLog.DebugLevel.TXT) }

func StatusLogDuration() lipgloss.TerminalColor { return GetColor(ThemeConfig.StatusLog.Duration.TXT) }

func OKStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.StatusLog.OKStatus.TXT))
}

func NOKStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.StatusLog.NOKStatus.TXT))
}

func WarnStatusMessage() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.StatusLog.WarnStatus.TXT))
}
