package style

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func TableStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		MarginLeft(2).
		Border(lipgloss.DoubleBorder(), true, true, true, true).
		BorderForeground(GetColor(ThemeConfig.Table.ActiveBorder))
}

func DimmedTableStyle() lipgloss.Style {
	return TableStyle().
		BorderForeground(GetColor(ThemeConfig.Table.DimmedBorder))
}

func tableHeaderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(GetColor(ThemeConfig.Table.ActiveHeader.BG))
}

func tableCellStyle() lipgloss.Style { return lipgloss.NewStyle() }

func tableSelectedStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(GetColor(ThemeConfig.Table.ActiveSelected.BG)).
		Foreground(GetColor(ThemeConfig.Table.ActiveSelected.TXT))
}

func dimmedTableHeaderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(GetColor(ThemeConfig.Table.DimmedHeader.BG)).
		Foreground(GetColor(ThemeConfig.Table.DimmedHeader.TXT))
}

func dimmedTableCellStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.Table.DimmedCell.TXT))
}

func dimmedTableSelectedStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(GetColor(ThemeConfig.Table.DimmedSelected.BG)).
		Foreground(GetColor(ThemeConfig.Table.DimmedSelected.TXT))
}

func PaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		MarginLeft(2).
		MarginBottom(1)
}

func DimmedPaginatorStyle() lipgloss.Style {
	return PaginatorStyle().Foreground(GetColor(ThemeConfig.Paginator.Dimmed))
}

func ActiveDotPaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(GetColor(ThemeConfig.Paginator.Active))
}

func InactiveDotPaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(GetColor(ThemeConfig.Paginator.Inactive))
}

func DimmedDotaginatorStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(GetColor(ThemeConfig.Paginator.Dimmed))
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
