package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/momarques/kibe/internal/ui/style"
)

const (
	paginatorDot = "•"
)

type paginatorModel struct {
	paginator.Model
}

func newPaginatorModel(itemsPerPage int) paginatorModel {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = itemsPerPage
	p.ActiveDot = style.ActiveDotPaginatorStyle().
		Render(paginatorDot)
	p.InactiveDot = style.InactiveDotPaginatorStyle().
		Render(paginatorDot)

	return paginatorModel{
		Model: p,
	}
}

func (p paginatorModel) view(dimm bool) string {
	if dimm {
		p.ActiveDot = style.DimmedDotaginatorStyle().Render(paginatorDot)
		p.InactiveDot = style.DimmedDotaginatorStyle().Render(paginatorDot)
		return style.DimmedPaginatorStyle().
			Render(p.View())
	}
	return style.PaginatorStyle().
		Render(p.View())
}
