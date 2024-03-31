package kube

import (
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
		List(c.Ctx, v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	p.pods = pods.Items
	return p, err
}

func (p Pod) Columns() (podAttributes []table.Column) {

	return append(podAttributes,
		table.Column{
			Title: "Name",
			Width: windowutil.ComputeWidthPercentage(nameColumnWidthPercentage)},
		table.Column{
			Title: "Ready",
			Width: windowutil.ComputeWidthPercentage(readyColumnWidthPercentage)},
		table.Column{
			Title: "Status",
			Width: windowutil.ComputeWidthPercentage(statusColumnWidthPercentage)},
		table.Column{
			Title: "Restarts",
			Width: windowutil.ComputeWidthPercentage(restartsColumnWidthPercentage)},
		table.Column{
			Title: "Node",
			Width: windowutil.ComputeWidthPercentage(nodeColumnWidthPercentage)},
		table.Column{
			Title: "Age",
			Width: windowutil.ComputeWidthPercentage(ageColumnWidthPercentage)},
	)
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
