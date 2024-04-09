package kube

import (
	"context"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespace struct {
	id, kind   string
	namespaces []corev1.Namespace
}

func NewNamespaceResource() Namespace { return Namespace{kind: "Namespace"} }

func (n Namespace) ID() string   { return n.id }
func (n Namespace) Kind() string { return n.kind }
func (n Namespace) SetID(id string) Resource {
	n.id = id
	return n
}

func (n Namespace) List(c ClientReady) (Resource, error) {
	namespaces, err := c.
		CoreV1().
		Namespaces().
		List(context.Background(), v1.ListOptions{})
	n.namespaces = namespaces.Items
	return n, err
}

func namespaceFieldWidth(width int, ns corev1.Namespace, _ int) int {
	if len(ns.Name) > width {
		return len(ns.Name)
	}
	return width
}

func (n Namespace) Columns() (namespaceAttributes []table.Column) {
	return append(namespaceAttributes,
		table.Column{Title: "Name", Width: lo.Reduce(n.namespaces, namespaceFieldWidth, 0)},
		table.Column{Title: "Age", Width: 20},
	)
}

func (n Namespace) Rows() (namespaceRows []table.Row) {
	for _, ns := range n.namespaces {
		namespaceRows = append(namespaceRows,
			table.Row{
				ns.Name,
				DeltaTime(ns.GetCreationTimestamp().Time, time.Now()),
			},
		)
	}
	return namespaceRows
}

type NamespaceSelected string

func (n NamespaceSelected) String() string { return string(n) }

type SelectNamespace []list.Item

func NewSelectNamespace(c ClientReady) func() tea.Msg {
	n, err := Namespace{}.List(c)
	if err != nil {
		c.Err <- err
	}
	return func() tea.Msg {
		return lo.Map(n.(Namespace).namespaces,
			func(item corev1.Namespace, _ int) list.Item {
				return NamespaceItem(item.Name)
			})
	}
}

type NamespaceItem string

func (ni NamespaceItem) Title() string       { return "Namespace: " + string(ni) }
func (ni NamespaceItem) FilterValue() string { return string(ni) }
func (ni NamespaceItem) Description() string { return "" }
