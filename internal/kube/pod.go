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
)

const (
	nameColumnWidthPercentage     = 34
	readyColumnWidthPercentage    = 4
	statusColumnWidthPercentage   = 10
	restartsColumnWidthPercentage = 6
	nodeColumnWidthPercentage     = 29
	ageColumnWidthPercentage      = 8
)

type Pod struct{ kind string }

func NewPodResource() *Pod  { return &Pod{kind: "Pod"} }
func (p *Pod) Kind() string { return p.kind }

func ListPods(c *ClientReady) []corev1.Pod {
	pods, err := c.Client.
		CoreV1().
		Pods(c.Namespace.NS).
		List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return pods.Items
}

func ListPodColumns(pods []corev1.Pod, width int) (podAttributes []table.Column) {

	return append(podAttributes,
		table.Column{Title: "Name", Width: computeWidthPercentage(
			width, nameColumnWidthPercentage)},
		table.Column{Title: "Ready", Width: computeWidthPercentage(
			width, readyColumnWidthPercentage)},
		table.Column{Title: "Status", Width: computeWidthPercentage(
			width, statusColumnWidthPercentage)},
		table.Column{Title: "Restarts", Width: computeWidthPercentage(
			width, restartsColumnWidthPercentage)},
		table.Column{Title: "Node", Width: computeWidthPercentage(
			width, nodeColumnWidthPercentage)},
		table.Column{Title: "Age", Width: computeWidthPercentage(
			width, ageColumnWidthPercentage)},
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
