package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

var loadInterval = 2 * time.Second

type contentState int

const (
	loaded contentState = iota
	notLoaded
	reload
)

type content struct {
	contentState

	client *kube.ClientReady
}

func newTableContent(c *kube.ClientReady) *content {
	return &content{
		contentState: notLoaded,
		client:       c,
	}
}

func (c *content) fetch(m table.Model) (table.Model, tea.Cmd) {
	columns, rows, title := FetchTableView(c.client)
	m.SetColumns(columns)
	m.SetRows(rows)
	m.SetHeight(
		windowutil.ComputePercentage(
			windowHeight, tableViewHeightPercentage))
	c.contentState = loaded
	return m, c.updateHeader(title, len(rows))
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
