package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	OKStatusMessage = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a2e3ad", Dark: "#a2e3ad"})
	NOKStatusMessage = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#f35143", Dark: "#f35143"})
	WarnStatusMessage = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#ffde6e", Dark: "#ffde6e"})

	StatusMessageStyle = OKStatusMessage.Render
)
