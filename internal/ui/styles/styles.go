package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	OKStatusMessage = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#a2e3ad", Dark: "#a2e3ad"})

	StatusMessageStyle = OKStatusMessage.Render

	UserStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9b5a46")).
			Background(lipgloss.NoColor{})

	NamespaceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#fcc493")).
			Background(lipgloss.NoColor{})
)

var (
	ViewTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#d65f50")).
			Padding(0, 1)

	ListActiveSelectionTitleStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("#d28169")).
					Foreground(lipgloss.Color("#ffffff")).
					Padding(0, 1, 0, 1).Bold(true)
	ListActiveSelectionDescStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("#9f614f")).
					Foreground(lipgloss.Color("#ffffff")).
					Padding(0, 1, 0, 1)
	ListDimmedDescStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#8d6f62")).
				Padding(0, 0, 0, 2)
	ListNormalTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#f8f3e8")).
				Padding(0, 0, 0, 2)

	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e4d491")).
			Background(lipgloss.NoColor{}).
			Padding(0, 0, 1, 0)

	ListFilterPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#b76dab", Dark: "#ECFD65"})

	ListFilterCursorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#a6ca83", Dark: "#ECFD65"})
)
