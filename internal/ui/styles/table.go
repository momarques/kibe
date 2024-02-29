package uistyles

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	TableStyle = lipgloss.NewStyle().
			MarginLeft(2).
			Border(lipgloss.DoubleBorder(), true, true, true, true).
			BorderForeground(lipgloss.Color("#ffb8bc"))
	DimmedTableStyle = TableStyle.
				Copy().
				BorderForeground(lipgloss.Color("#616161"))

	tableHeaderStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#c5636a"))
	tableCellStyle     = lipgloss.NewStyle()
	tableSelectedStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#ffb1b5")).
				Foreground(lipgloss.Color("#322223"))

	dimmedTableHeaderStyle = lipgloss.NewStyle().
				Background(lipgloss.AdaptiveColor{Light: "#494949", Dark: "#494949"}).
				Foreground(lipgloss.Color("#616161"))
	dimmedTableCellStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#616161"))

	dimmedTableSelectedStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("#616161")).
					Foreground(lipgloss.Color("#222222"))

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
