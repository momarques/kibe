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
)

type Service struct {
	kind     string
	services []corev1.Service
}

func NewServiceResource() Service { return Service{kind: "Service"} }
func (s Service) Kind() string    { return s.kind }

func (s Service) List(c *ClientReady) Resource {
	services, err := c.
		CoreV1().
		Services(string(c.NamespaceSelected)).
		List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	s.services = services.Items
	return s
}
func (s Service) Columns() (serviceAttributes []table.Column) {
	return append(serviceAttributes,
		table.Column{Title: "Name", Width: serviceFieldWidth("Name", s.services)},
		table.Column{Title: "Type", Width: serviceFieldWidth("Type", s.services)},
		table.Column{Title: "ClusterIP", Width: serviceFieldWidth("ClusterIP", s.services)},
		table.Column{Title: "ExternalIP", Width: serviceFieldWidth("ExternalIP", s.services)},
		table.Column{Title: "Ports", Width: serviceFieldWidth("Ports", s.services)},
		table.Column{Title: "Age", Width: 20},
	)
}

func (s Service) Rows() (serviceRows []table.Row) {
	for _, svc := range s.services {

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

	return lo.Reduce(services,
		func(width int, svc corev1.Service, _ int) int {

			switch fieldName {
			case "Name":
				fieldValue = svc.Name
			case "Type":
				fieldValue = string(svc.Spec.Type)
			case "ClusterIP":
				fieldValue = svc.Spec.ClusterIP
			case "ExternalIP":
				fieldValue = strings.Join(
					svc.Spec.ExternalIPs, ", ")

				// workaround so the column can have sufficient space to print column name, in case the value is empty
				if len(fieldValue) < 1 {
					fieldValue = "ExternalIP"
				}
			case "Ports":
				fieldValue = servicePortsAsString(svc.Spec.Ports)
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
		portAsString := fmt.Sprintf("%s::%d", port.Name, port.Port)
		if port.NodePort != 0 {
			portAsString = fmt.Sprintf("%s::%d", portAsString, port.NodePort)
		}

		ports = append(ports, portAsString)
	}
	return strings.Join(ports, ", ")
}
