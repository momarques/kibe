package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	TableStyle = lipgloss.NewStyle().
			MarginLeft(2).
			Border(lipgloss.DoubleBorder(), true, true, true, true).
			BorderForeground(lipgloss.Color("#ffb8bc"))

	TableHeaderStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#c5636a"))
	TableCellStyle = lipgloss.NewStyle()

	TableSelectedStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#ffb1b5")).
				Foreground(lipgloss.Color("#322223"))

	PaginatorStyle = lipgloss.NewStyle().MarginLeft(2)
)
