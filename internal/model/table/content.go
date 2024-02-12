package tablemodel

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/kube/pod"
	"k8s.io/client-go/kubernetes"
)

func FetchTable(kind, namespace string, client *kubernetes.Clientset) ([]table.Column, []table.Row) {
	switch kind {
	case "Pod":
		pods := pod.FetchResources(namespace, client)
		podColumns := pod.FetchColumns(pods)

		return podColumns, pod.RetrievePodListAsTableRows(pods)
	}
	return nil, nil
}
