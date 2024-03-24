package style

import "github.com/charmbracelet/lipgloss"

var (
	OKStatusMessage = lipgloss.NewStyle().
			Foreground(GetColor(defaultThemeConfig.StatusLog.OKStatus.TXT))
	NOKStatusMessage = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.StatusLog.NOKStatus.TXT))
	WarnStatusMessage = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.StatusLog.WarnStatus.TXT))
)
