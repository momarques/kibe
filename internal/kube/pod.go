package kube

import (
	"context"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Pod struct{ kind string }

func NewPodResource() *Pod  { return &Pod{kind: "Pod"} }
func (p *Pod) Kind() string { return p.kind }

func ListPods(namespace string, client *kubernetes.Clientset) []corev1.Pod {
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return pods.Items
}

func ListPodColumns(pods []corev1.Pod) (podAttributes []table.Column) {
	return append(podAttributes,
		table.Column{Title: "Name", Width: podFieldWidth("Name", pods)},
		table.Column{Title: "Ready", Width: 10},
		table.Column{Title: "Status", Width: 20},
		table.Column{Title: "Restarts", Width: 10},
		table.Column{Title: "Node", Width: podFieldWidth("Node", pods)},
		table.Column{Title: "Age", Width: 20},
	)
}

func RetrievePodListAsTableRows(pods []corev1.Pod) (podRows []table.Row) {
	for _, pod := range pods {
		podRows = append(podRows,
			table.Row{
				pod.Name,
				checkReadyContainers(
					pod.Status.ContainerStatuses),
				string(
					pod.Status.Phase),
				checkRestartedContainers(
					pod.Status.ContainerStatuses),
				pod.Spec.NodeName,
				DeltaTime(
					pod.GetCreationTimestamp().Time),
			},
		)
	}
	return podRows
}

func podFieldWidth(fieldName string, pods []corev1.Pod) int {
	var fieldValue string
	return lo.Reduce(pods, func(width int, pod corev1.Pod, _ int) int {

		switch fieldName {
		case "Name":
			fieldValue = pod.Name
		case "Node":
			fieldValue = pod.Spec.NodeName
		}

		if len(fieldValue) > width {
			return len(fieldValue)
		}
		return width
	}, 0)
}

func checkReadyContainers(containers []corev1.ContainerStatus) string {
	return fmt.Sprintf("%d/%d",
		lo.CountBy(containers, func(c corev1.ContainerStatus) bool {
			return c.Ready
		}),
		len(containers))
}

func checkRestartedContainers(containers []corev1.ContainerStatus) string {
	return strconv.Itoa(
		lo.Reduce(containers, func(restarts int, container corev1.ContainerStatus, _ int) int {
			return restarts + int(container.RestartCount)
		}, 0))
}
