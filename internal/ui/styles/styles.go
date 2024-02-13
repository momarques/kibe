package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	UserStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e9b7da")).
			Background(lipgloss.NoColor{})

	NamespaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a2e3ad")).
			Background(lipgloss.NoColor{})
)
