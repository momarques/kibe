package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
)

// var loadInterval = 2 * time.Second

type contentState int

const (
	loaded contentState = iota
	notLoaded
	reload
)

type content struct {
	contentState

	client *kube.ClientReady

	rows      []table.Row
	paginator paginator.Model
}

func newTableContent(c *kube.ClientReady, paginator paginator.Model) *content {
	return &content{
		contentState: notLoaded,
		client:       c,
		paginator:    paginator,
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

func (c *content) fetchTableItems(m table.Model) (table.Model, tea.Cmd) {
	columns, rows, title := FetchTableView(c.client)

	m.SetColumns(columns)

	c.paginator.SetTotalPages(len(rows))
	c.rows = rows
	return m, c.updateHeader(title, len(rows))
}

func (c *content) fetchPageItems(m table.Model) table.Model {
	start, end := c.paginator.GetSliceBounds(len(c.rows))

	rows := c.rows[start:end]
	m.SetRows(rows)
	return m
}
