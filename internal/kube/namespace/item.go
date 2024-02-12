package namespace

import (
	"context"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct{ kind string }

func New() *Namespace             { return &Namespace{kind: "Namespace"} }
func (p *Namespace) Kind() string { return p.kind }

func FetchResources(client *kubernetes.Clientset) []corev1.Namespace {
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return namespaces.Items
}

func RetrieveNamespaceListAsTableRows(namespaces []corev1.Namespace) (namespaceRows []table.Row) {
	for _, ns := range namespaces {
		namespaceRows = append(namespaceRows,
			table.Row{
				ns.Name,
				kube.DeltaTime(
					ns.GetCreationTimestamp().Time),
			},
		)
	}
	return namespaceRows
}

func FetchColumns(namespaces []corev1.Namespace) (namespaceAttributes []table.Column) {
	return append(namespaceAttributes,
		table.Column{Title: "Name", Width: namespaceFieldWidth("Name", namespaces)},
		table.Column{Title: "Age", Width: 20},
	)
}

func namespaceFieldWidth(fieldName string, namespaces []corev1.Namespace) int {
	return lo.Reduce(namespaces, func(width int, ns corev1.Namespace, _ int) int {
		if len(ns.Name) > width {
			return len(ns.Name)
		}
		return width
	}, 0)
}
