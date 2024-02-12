package kube

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Service struct{ kind string }

func NewServiceResource() *Service { return &Service{kind: "Service"} }
func (s *Service) Kind() string    { return s.kind }

func ListServices(namespace string, client *kubernetes.Clientset) []corev1.Service {
	services, err := client.CoreV1().Services(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return services.Items
}
func ListServiceColumns(services []corev1.Service) (serviceAttributes []table.Column) {
	return append(serviceAttributes,
		table.Column{Title: "Name", Width: serviceFieldWidth("Name", services)},
		table.Column{Title: "Type", Width: serviceFieldWidth("Name", services)},
		table.Column{Title: "ClusterIP", Width: serviceFieldWidth("Name", services)},
		table.Column{Title: "ExternalIP", Width: 100},
		table.Column{Title: "Ports", Width: 100},
		table.Column{Title: "Age", Width: 20},
	)
}

func RetrieveServiceListAsTableRows(services []corev1.Service) (serviceRows []table.Row) {
	for _, svc := range services {

		serviceRows = append(serviceRows,
			table.Row{
				svc.Name,
				string(
					svc.Spec.Type),
				svc.Spec.ClusterIP,
				strings.Join(
					svc.Spec.ExternalIPs, ", "),
				servicePortsAsString(
					svc.Spec.Ports),
				DeltaTime(
					svc.GetCreationTimestamp().Time),
			},
		)
	}
	return serviceRows
}

func serviceFieldWidth(fieldName string, services []corev1.Service) int {
	var fieldValue string

	return lo.Reduce(services, func(width int, svc corev1.Service, _ int) int {

		switch fieldName {
		case "Name":
			fieldValue = svc.Name
		case "Type":
			fieldValue = string(svc.Spec.Type)
		case "ClusterIP":
			fieldValue = svc.Spec.ClusterIP
		}

		if len(fieldValue) > width {
			return len(fieldValue)
		}
		return width
	}, 0)
}

func servicePortsAsString(services []corev1.ServicePort) string {
	var ports []string

	for _, port := range services {
		ports = append(ports,
			fmt.Sprintf("%s::%s::%s", port.Name, string(port.Port), string(port.NodePort)))
	}
	return strings.Join(ports, ", ")
}
