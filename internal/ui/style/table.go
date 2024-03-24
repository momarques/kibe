package style

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	TableStyle = lipgloss.NewStyle().
			MarginLeft(2).
			Border(lipgloss.DoubleBorder(), true, true, true, true).
			BorderForeground(GetColor(defaultThemeConfig.Table.ActiveBorder))
	DimmedTableStyle = TableStyle.
				Copy().
				BorderForeground(GetColor(defaultThemeConfig.Table.DimmedBorder))

	tableHeaderStyle = lipgloss.NewStyle().
				Background(GetColor(defaultThemeConfig.Table.ActiveHeader.BG))
	tableCellStyle     = lipgloss.NewStyle()
	tableSelectedStyle = lipgloss.NewStyle().
				Background(GetColor(defaultThemeConfig.Table.ActiveSelected.BG)).
				Foreground(GetColor(defaultThemeConfig.Table.ActiveSelected.TXT))

	dimmedTableHeaderStyle = lipgloss.NewStyle().
				Background(GetColor(defaultThemeConfig.Table.DimmedHeader.BG)).
				Foreground(GetColor(defaultThemeConfig.Table.DimmedHeader.TXT))
	dimmedTableCellStyle = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.Table.DimmedCell.TXT))

	dimmedTableSelectedStyle = lipgloss.NewStyle().
					Background(GetColor(defaultThemeConfig.Table.DimmedSelected.BG)).
					Foreground(GetColor(defaultThemeConfig.Table.DimmedSelected.TXT))

	PaginatorStyle = lipgloss.NewStyle().MarginLeft(2)
)

func NewTableStyle(dimm bool) table.Styles {
	s := table.DefaultStyles()

	s.Cell = tableCellStyle
	s.Header = tableHeaderStyle
	s.Selected = tableSelectedStyle

	if dimm {
		s.Cell = dimmedTableCellStyle
		s.Header = dimmedTableHeaderStyle
		s.Selected = dimmedTableSelectedStyle
	}

	return s
}
