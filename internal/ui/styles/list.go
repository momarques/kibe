package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

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

	ListFilterPromptStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#b76dab", Dark: "#ECFD65"})

	ListFilterCursorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#a6ca83", Dark: "#ECFD65"})
)
