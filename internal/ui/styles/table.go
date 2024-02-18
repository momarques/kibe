package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	TableStyle = lipgloss.NewStyle().
			MarginLeft(2).
			Border(lipgloss.DoubleBorder(), true, true, true, true).
			BorderForeground(lipgloss.Color("#ffb8bc"))

	TableHeaderStyle = lipgloss.NewStyle().
				Blink(false).
				Background(lipgloss.Color("#c5636a"))
	TableSelectedStyle = lipgloss.NewStyle().
				Blink(false).
				Background(lipgloss.Color("#ffb1b5")).
				Foreground(lipgloss.Color("#322223"))
)
