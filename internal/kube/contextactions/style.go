package contextactions

import "github.com/charmbracelet/lipgloss"

var (
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	userStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e9b7da")).
			Background(lipgloss.NoColor{})

	namespaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a2e3ad")).
			Background(lipgloss.NoColor{})
)
