package style

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

func TableStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		MarginLeft(2).
		Border(lipgloss.DoubleBorder(), true, true, true, true).
		BorderForeground(theme.GetColor(theme.Selected.Table.ActiveBorder))
}

func DimmedTableStyle() lipgloss.Style {
	return TableStyle().
		BorderForeground(theme.GetColor(theme.Selected.Table.DimmedBorder))
}

func tableHeaderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.GetColor(theme.Selected.Table.ActiveHeader.BG))
}

func tableCellStyle() lipgloss.Style { return lipgloss.NewStyle() }

func tableSelectedStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.GetColor(theme.Selected.Table.ActiveSelected.BG)).
		Foreground(theme.GetColor(theme.Selected.Table.ActiveSelected.TXT))
}

func dimmedTableHeaderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.GetColor(theme.Selected.Table.DimmedHeader.BG)).
		Foreground(theme.GetColor(theme.Selected.Table.DimmedHeader.TXT))
}

func dimmedTableCellStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.Table.DimmedCell.TXT))
}

func dimmedTableSelectedStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.GetColor(theme.Selected.Table.DimmedSelected.BG)).
		Foreground(theme.GetColor(theme.Selected.Table.DimmedSelected.TXT))
}

func PaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		MarginLeft(2).
		MarginBottom(1)
}

func DimmedPaginatorStyle() lipgloss.Style {
	return PaginatorStyle().Foreground(theme.GetColor(theme.Selected.Paginator.Dimmed))
}

func ActiveDotPaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.GetColor(theme.Selected.Paginator.Active))
}

func InactiveDotPaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.GetColor(theme.Selected.Paginator.Inactive))
}

func DimmedDotaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.GetColor(theme.Selected.Paginator.Dimmed))
}

func NewTableStyle(dimm bool) table.Styles {
	s := table.DefaultStyles()

	s.Cell = tableCellStyle()
	s.Header = tableHeaderStyle()
	s.Selected = tableSelectedStyle()

	if dimm {
		s.Cell = dimmedTableCellStyle()
		s.Header = dimmedTableHeaderStyle()
		s.Selected = dimmedTableSelectedStyle()
	}
	return s
}
