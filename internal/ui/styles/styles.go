package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	OKStatusMessage = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a2e3ad", Dark: "#a2e3ad"})

	StatusMessageStyle = OKStatusMessage.Render

	UserStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e9b7da")).
			Background(lipgloss.NoColor{})

	NamespaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a2e3ad")).
			Background(lipgloss.NoColor{})
)
