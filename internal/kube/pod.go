package kube

import (
	"context"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	nameColumnWidthPercentage     int = 39
	readyColumnWidthPercentage    int = 5
	statusColumnWidthPercentage   int = 11
	restartsColumnWidthPercentage int = 6
	nodeColumnWidthPercentage     int = 29
	ageColumnWidthPercentage      int = 8
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

func ListPodColumns(pods []corev1.Pod) (podAttributes []table.Column) {
	var (
		nameColumnWidth     int = windowutil.ComputeWidthPercentage(nameColumnWidthPercentage)
		readyColumnWidth    int = windowutil.ComputeWidthPercentage(readyColumnWidthPercentage)
		statusColumnWidth   int = windowutil.ComputeWidthPercentage(statusColumnWidthPercentage)
		restartsColumnWidth int = windowutil.ComputeWidthPercentage(restartsColumnWidthPercentage)
		nodeColumnWidth     int = windowutil.ComputeWidthPercentage(nodeColumnWidthPercentage)
		ageColumnWidth      int = windowutil.ComputeWidthPercentage(ageColumnWidthPercentage)
	)

	logging.Log.Info(
		nameColumnWidth,
		readyColumnWidth,
		statusColumnWidth,
		restartsColumnWidth,
		nodeColumnWidth,
		ageColumnWidth,
	)
	return append(podAttributes,
		table.Column{Title: "Name", Width: nameColumnWidth},
		table.Column{Title: "Ready", Width: readyColumnWidth},
		table.Column{Title: "Status", Width: statusColumnWidth},
		table.Column{Title: "Restarts", Width: restartsColumnWidth},
		table.Column{Title: "Node", Width: nodeColumnWidth},
		table.Column{Title: "Age", Width: ageColumnWidth},
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
		lo.CountBy(containers,
			func(c corev1.ContainerStatus) bool {
				return c.Ready
			}),
		len(containers))
}

func checkRestartedContainers(containers []corev1.ContainerStatus) string {
	return strconv.Itoa(
		lo.Reduce(containers,
			func(restarts int, container corev1.ContainerStatus, _ int) int {
				return restarts + int(container.RestartCount)
			}, 0))
}
