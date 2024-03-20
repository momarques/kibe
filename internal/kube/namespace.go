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

type Namespace struct {
	kind       string
	namespaces []corev1.Namespace
}

func NewNamespaceResource() Namespace { return Namespace{kind: "Namespace"} }
func (n Namespace) Kind() string      { return n.kind }

func (n Namespace) List(c *ClientReady) Resource {
	namespaces, err := c.
		CoreV1().
		Namespaces().
		List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	n.namespaces = namespaces.Items
	return n
}

func (n Namespace) Columns() (namespaceAttributes []table.Column) {
	return append(namespaceAttributes,
		table.Column{Title: "Name", Width: n.namespaceFieldWidth()},
		table.Column{Title: "Age", Width: 20},
	)
}

func (n Namespace) Rows() (namespaceRows []table.Row) {
	for _, ns := range n.namespaces {
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

func (n Namespace) namespaceFieldWidth() int {
	return lo.Reduce(n.namespaces,
		func(width int, ns corev1.Namespace, _ int) int {
			if len(ns.Name) > width {
				return len(ns.Name)
			}
			return width
		}, 0)
}

type SelectNamespace []list.Item
type NamespaceSelected string

func NewSelectNamespace(c *ClientReady) func() tea.Msg {
	n := Namespace{}.List(c)
	return func() tea.Msg {
		return SelectNamespace(n.(Namespace).newNamespaceList())
	}
}

type NamespaceItem string

func (ni NamespaceItem) Title() string       { return "Namespace: " + string(ni) }
func (ni NamespaceItem) FilterValue() string { return string(ni) }
func (ni NamespaceItem) Description() string { return "" }

func (n Namespace) newNamespaceList() []list.Item {
	namespaceList := []list.Item{}

	for _, ns := range n.namespaces {
		namespaceList = append(namespaceList, NamespaceItem(ns.Name))
	}
	return namespaceList
}
