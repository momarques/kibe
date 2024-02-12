package tablemodel

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/kube"
	"k8s.io/client-go/kubernetes"
)

func FetchTable(resourceKind, resourceNamespace string, client *kubernetes.Clientset) ([]table.Column, []table.Row) {
	switch resourceKind {
	case "Pod":
		pods := kube.ListPods(resourceNamespace, client)
		podColumns := kube.ListPodColumns(pods)

		return podColumns, kube.RetrievePodListAsTableRows(pods)
	case "Namespace":
		ns := kube.ListNamespaces(client)
		nsColumns := kube.ListNamespaceColumns(ns)

		return nsColumns, kube.RetrieveNamespaceListAsTableRows(ns)
	case "Service":
		svc := kube.ListServices(resourceNamespace, client)
		svcColumns := kube.ListServiceColumns(svc)

		return svcColumns, kube.RetrieveServiceListAsTableRows(svc)
	}
	return nil, nil
}
