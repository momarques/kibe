package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
)

type tableContent struct {
	syncState

	client *kube.ClientReady

	columns        []table.Column
	rows           []table.Row
	paginatorModel paginator.Model
}

func newTableContent(c *kube.ClientReady) *tableContent {
	return &tableContent{
		syncState:      unsynced,
		client:         c,
		paginatorModel: newPaginatorModel(),
	}
}

func (c *tableContent) applyTableItems(m table.Model) (table.Model, tea.Cmd) {
	m.SetColumns(c.columns)

	start, end := c.paginatorModel.GetSliceBounds(len(c.rows))
	m.SetRows(c.rows[start:end])
	return m, c.updateHeader(len(c.rows))
}
