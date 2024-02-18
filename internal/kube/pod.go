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
		table.Column{Title: "Name", Width: windowutil.ComputePercentage(
			width, nameColumnWidthPercentage)},
		table.Column{Title: "Ready", Width: windowutil.ComputePercentage(
			width, readyColumnWidthPercentage)},
		table.Column{Title: "Status", Width: windowutil.ComputePercentage(
			width, statusColumnWidthPercentage)},
		table.Column{Title: "Restarts", Width: windowutil.ComputePercentage(
			width, restartsColumnWidthPercentage)},
		table.Column{Title: "Node", Width: windowutil.ComputePercentage(
			width, nodeColumnWidthPercentage)},
		table.Column{Title: "Age", Width: windowutil.ComputePercentage(
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

func DescribePod(c *ClientReady, podID string) *corev1.Pod {
	pod, err := c.Client.
		CoreV1().
		Pods(c.Namespace.NS).
		Get(context.Background(), podID, v1.GetOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return pod
}

func NewPodDescription(c *ClientReady, podID string) PodDescription {
	pod := DescribePod(c, podID)

	return PodDescription{
		Overview: newPodOverview(pod),
	}
}
