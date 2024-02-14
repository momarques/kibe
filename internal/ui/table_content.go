package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/kube"
)

type contentState int

const (
	loaded contentState = iota
	notLoaded
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

func (c *content) fetch(m *table.Model) *table.Model {
	columns, rows := FetchTable(c.client)
	m.SetColumns(columns)
	m.SetRows(rows)
	c.contentState = loaded
	return m
}

func (c *content) setTableTitle(m *table.Model) {

}

func (c *content) setTableResourceLegend() {

}

func FetchTable(c *kube.ClientReady) ([]table.Column, []table.Row) {
	switch c.ResourceSelected.R.(type) {
	case *kube.Pod:
		pods := kube.ListPods(c)
		podColumns := kube.ListPodColumns(pods)

		return podColumns, kube.RetrievePodListAsTableRows(pods)
	case *kube.Namespace:
		ns := kube.ListNamespaces(c)
		nsColumns := kube.ListNamespaceColumns(ns)

		return nsColumns, kube.RetrieveNamespaceListAsTableRows(ns)
	case *kube.Service:
		svc := kube.ListServices(c)
		svcColumns := kube.ListServiceColumns(svc)

		return svcColumns, kube.RetrieveServiceListAsTableRows(svc)
	}
	return nil, nil
}