package uistyles

import "github.com/charmbracelet/lipgloss"

var (
	TableStyle = lipgloss.NewStyle().
		MarginLeft(2).
		Border(lipgloss.DoubleBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("#ffb8bc"))
)
