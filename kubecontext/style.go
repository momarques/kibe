package kubecontext

import "github.com/charmbracelet/lipgloss"

var (
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	userStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffeaaf")).
			Background(lipgloss.NoColor{})

	namespaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffaeca")).
			Background(lipgloss.NoColor{})
)
