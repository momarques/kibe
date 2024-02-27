package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
)

func newPaginatorUI() paginator.Model {

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 16
	p.ActiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).
		MarginLeft(2).
		Render("•")
	p.InactiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).
		MarginLeft(2).
		Render("•")

	return p
}
