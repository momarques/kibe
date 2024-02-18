package uistyles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ViewTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#d65f50")).
			Padding(0, 1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e4d491")).
			Background(lipgloss.NoColor{}).
			Padding(0, 0, 1, 0)
)
