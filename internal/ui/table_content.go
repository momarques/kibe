package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/kube"
)

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
