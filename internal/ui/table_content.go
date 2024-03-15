package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type tableContent struct {
	syncState

	columns   []table.Column
	rows      []table.Row
	paginator paginator.Model
}

func newTableContent() tableContent {
	return tableContent{
		syncState: unsynced,
		paginator: newPaginatorModel(15),
	}
}

func (m tableModel) applyTableItems() (tableModel, tea.Cmd) {
	m.SetColumns(m.columns)

	start, end := m.paginator.GetSliceBounds(len(m.rows))
	m.SetRows(m.rows[start:end])
	return m, m.updateHeader(len(m.rows))
}
