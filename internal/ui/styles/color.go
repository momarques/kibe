package uistyles

import (
	"github.com/charmbracelet/lipgloss"
)

func ColorizeTabKeys(row, col int) lipgloss.Style {
	switch {
	case col == 0:
		return resourceSectionDescriptionStyle.
			Copy().
			Background(lipgloss.NoColor{}).
			Foreground(lipgloss.Color("#ff9184")).
			Padding(0, 1).
			Bold(true)
	}
	return lipgloss.NewStyle()
}
