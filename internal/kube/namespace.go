package kube

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespace struct{ kind string }

func NewNamespaceResource() *Namespace { return &Namespace{kind: "Namespace"} }
func (n *Namespace) Kind() string      { return n.kind }

func ListNamespaces(c *ClientReady) []corev1.Namespace {
	namespaces, err := c.Client.
		CoreV1().
		Namespaces().
		List(context.Background(), v1.ListOptions{})
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
	return lo.Reduce(namespaces,
		func(width int, ns corev1.Namespace, _ int) int {
			if len(ns.Name) > width {
				return len(ns.Name)
			}
			return width
		}, 0)
}

type SelectNamespace struct{ Namespaces []list.Item }
type NamespaceSelected struct{ NS string }

func NewSelectNamespace(c *ClientReady) func() tea.Msg {
	return func() tea.Msg {
		return SelectNamespace{
			Namespaces: newNamespaceList(c)}
	}
}

type NamespaceItem struct{ name string }

func (ni NamespaceItem) Title() string       { return "Namespace: " + ni.name }
func (ni NamespaceItem) FilterValue() string { return ni.name }
func (ni NamespaceItem) Description() string { return "" }

func newNamespaceList(c *ClientReady) []list.Item {
	namespaces := ListNamespaces(c)

	namespaceList := []list.Item{}

	for _, ns := range namespaces {
		namespaceList = append(namespaceList, NamespaceItem{
			name: ns.Name,
		})
	}
	return namespaceList
}

func NamespacesAsList(c *ClientReady) ([]list.Item, error) {
	return newNamespaceList(c), nil
}
