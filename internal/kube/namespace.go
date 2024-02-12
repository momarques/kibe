package kube

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct{ kind string }

func NewNamespaceResource() *Namespace { return &Namespace{kind: "Namespace"} }
func (n *Namespace) Kind() string      { return n.kind }

func ListNamespaces(client *kubernetes.Clientset) []corev1.Namespace {
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return namespaces.Items
}

func ListNamespaceColumns(namespaces []corev1.Namespace) (namespaceAttributes []table.Column) {
	return append(namespaceAttributes,
		table.Column{Title: "Name", Width: namespaceFieldWidth("Name", namespaces)},
		table.Column{Title: "Age", Width: 20},
	)
}

func RetrieveNamespaceListAsTableRows(namespaces []corev1.Namespace) (namespaceRows []table.Row) {
	for _, ns := range namespaces {
		namespaceRows = append(namespaceRows,
			table.Row{
				ns.Name,
				DeltaTime(
					ns.GetCreationTimestamp().Time),
			},
		)
	}
	return namespaceRows
}

func namespaceFieldWidth(fieldName string, namespaces []corev1.Namespace) int {
	return lo.Reduce(namespaces, func(width int, ns corev1.Namespace, _ int) int {
		if len(ns.Name) > width {
			return len(ns.Name)
		}
		return width
	}, 0)
}

type NamespaceItem struct{ corev1.Namespace }

func (ni NamespaceItem) Title() string       { return "Namespace: " + ni.Name }
func (ni NamespaceItem) FilterValue() string { return ni.Name }
func (ni NamespaceItem) Description() string { return "" }

func newNamespaceList(client *kubernetes.Clientset) []list.Item {
	namespaces := ListNamespaces(client)

	namespaceList := []list.Item{}

	for _, ns := range namespaces {
		namespaceList = append(namespaceList, NamespaceItem{
			Namespace: ns,
		})
	}
	return namespaceList
}

func NamespacesAsList(context string, client *kubernetes.Clientset) ([]list.Item, error) {
	return newNamespaceList(client), nil
}
