package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type tableContent struct {
	syncState

	columns        []table.Column
	rows           []table.Row
	paginatorModel paginator.Model
}

func newTableContent() *tableContent {
	return &tableContent{
		syncState:      unsynced,
		paginatorModel: newPaginatorModel(),
	}
}

func (c *tableContent) applyTableItems(m table.Model) (table.Model, tea.Cmd) {
	m.SetColumns(c.columns)

	start, end := c.paginatorModel.GetSliceBounds(len(c.rows))
	m.SetRows(c.rows[start:end])
	return m, c.updateHeader(len(c.rows))
}
