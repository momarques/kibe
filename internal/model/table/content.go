package tablemodel

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/kube/namespace"
	"github.com/momarques/kibe/internal/kube/pod"
	"k8s.io/client-go/kubernetes"
)

func FetchTable(resourceKind, resourceNamespace string, client *kubernetes.Clientset) ([]table.Column, []table.Row) {
	switch resourceKind {
	case "Pod":
		pods := pod.FetchResources(resourceNamespace, client)
		podColumns := pod.FetchColumns(pods)

		return podColumns, pod.RetrievePodListAsTableRows(pods)
	case "Namespace":
		ns := namespace.FetchResources(client)
		nsColumns := namespace.FetchColumns(ns)

		return nsColumns, namespace.RetrieveNamespaceListAsTableRows(ns)
	}
	return nil, nil
}
