package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

func newPaginatorUI() paginator.Model {

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 16
	p.ActiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "#ffb1b5"}).
		Render("•")
	p.InactiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "#624548"}).
		Render("•")

	return p
}

func (m CoreUI) viewPaginatorUI() string {
	return uistyles.PaginatorStyle.
		Copy().
		MarginRight(35).
		Render(m.paginatorModel.View())
}
