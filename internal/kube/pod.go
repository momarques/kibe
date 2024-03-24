package kube

import (
	"context"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
	"github.com/momarques/kibe/internal/ui/style"
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

type Pod struct {
	id, kind string
	pods     []corev1.Pod
}

func NewPodResource() Pod { return Pod{kind: "Pod"} }

func (p Pod) ID() string   { return p.id }
func (p Pod) Kind() string { return p.kind }
func (p Pod) SetID(id string) Resource {
	p.id = id
	return p
}

func (p Pod) List(c *ClientReady) (Resource, error) {
	pods, err := c.
		CoreV1().
		Pods(c.NamespaceSelected.String()).
		List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	p.pods = pods.Items
	return p, err
}

var (
	nameColumnWidth     int = windowutil.ComputeWidthPercentage(nameColumnWidthPercentage)
	readyColumnWidth    int = windowutil.ComputeWidthPercentage(readyColumnWidthPercentage)
	statusColumnWidth   int = windowutil.ComputeWidthPercentage(statusColumnWidthPercentage)
	restartsColumnWidth int = windowutil.ComputeWidthPercentage(restartsColumnWidthPercentage)
	nodeColumnWidth     int = windowutil.ComputeWidthPercentage(nodeColumnWidthPercentage)
	ageColumnWidth      int = windowutil.ComputeWidthPercentage(ageColumnWidthPercentage)
)

func (p Pod) Columns() (podAttributes []table.Column) {

	return append(podAttributes,
		table.Column{Title: "Name", Width: nameColumnWidth},
		table.Column{Title: "Ready", Width: readyColumnWidth},
		table.Column{Title: "Status", Width: statusColumnWidth},
		table.Column{Title: "Restarts", Width: restartsColumnWidth},
		table.Column{Title: "Node", Width: nodeColumnWidth},
		table.Column{Title: "Age", Width: ageColumnWidth},
	)
}

func podPhaseColor(p corev1.PodPhase) string {
	pString := string(p)

	switch p {
	case corev1.PodRunning, corev1.PodSucceeded:
		return style.OKStatusMessage().
			Render(pString)
	case corev1.PodPending:
		return style.WarnStatusMessage().Render(pString)
	case corev1.PodFailed:
		return style.NOKStatusMessage().Render(pString)
	case corev1.PodUnknown:
		return style.NoneStatusMessage().Render(pString)
	}

	return pString
}

func (p Pod) Rows() (podRows []table.Row) {
	for _, pod := range p.pods {
		podRows = append(podRows,
			table.Row{
				pod.Name,
				checkReadyContainers(pod.Status.ContainerStatuses),
				string(pod.Status.Phase),
				checkRestartedContainers(pod.Status.ContainerStatuses),
				pod.Spec.NodeName,
				DeltaTime(pod.GetCreationTimestamp().Time),
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
