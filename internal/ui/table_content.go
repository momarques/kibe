package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
)

var loadInterval = 2 * time.Second

type tableContent struct {
	syncState

	client *kube.ClientReady

	rows           []table.Row
	paginatorModel paginator.Model
}

func newTableContent(c *kube.ClientReady) *tableContent {
	return &tableContent{
		syncState:      unsynced,
		client:         c,
		paginatorModel: newPaginatorUI(),
	}
}

func FetchTableView(c *kube.ClientReady) ([]table.Column, []table.Row, string) {
	switch c.ResourceSelected.R.(type) {
	case *kube.Pod:
		pods := kube.ListPods(c)
		podColumns := kube.ListPodColumns(pods, windowWidth)

		return podColumns, kube.RetrievePodListAsTableRows(pods), "Pod interaction"
	case *kube.Namespace:
		ns := kube.ListNamespaces(c)
		nsColumns := kube.ListNamespaceColumns(ns)

		return nsColumns, kube.RetrieveNamespaceListAsTableRows(ns), "Namespace interaction"
	case *kube.Service:
		svc := kube.ListServices(c)
		svcColumns := kube.ListServiceColumns(svc)

		return svcColumns, kube.RetrieveServiceListAsTableRows(svc), "Service interaction"
	}
	return nil, nil, ""
}

func (c *tableContent) fetchTableItems(m table.Model) (table.Model, tea.Cmd) {
	columns, rows, title := FetchTableView(c.client)

	m.SetColumns(columns)

	c.paginatorModel.SetTotalPages(len(rows))
	c.rows = rows
	return m, c.updateHeader(title, len(rows))
}

func (c *tableContent) fetchPageItems(m table.Model) table.Model {
	start, end := c.paginatorModel.GetSliceBounds(len(c.rows))

	rows := c.rows[start:end]
	m.SetRows(rows)
	return m
}
