package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	OKStatusMessage = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a2e3ad", Dark: "#a2e3ad"})

	StatusMessageStyle = OKStatusMessage.Render
)
