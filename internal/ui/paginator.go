package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type paginatorModel struct {
	paginator.Model
}

func newPaginatorModel(itemsPerPage int) paginatorModel {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = itemsPerPage
	p.ActiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "#ffb1b5"}).
		Render("•")
	p.InactiveDot = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "#624548"}).
		Render("•")

	return paginatorModel{
		Model: p,
	}
}

func (p paginatorModel) view() string {
	return uistyles.PaginatorStyle.
		Copy().
		MarginRight(40).
		MarginBottom(1).
		Render(p.View())
}
