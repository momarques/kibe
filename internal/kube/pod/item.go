package pod

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	"github.com/thoas/go-funk"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Pod struct{ kind string }

func New() *Pod             { return &Pod{kind: "Pod"} }
func (p *Pod) Kind() string { return p.kind }

func FetchResources(namespace string, client *kubernetes.Clientset) []corev1.Pod {
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return pods.Items
}

func RetrievePodListAsTableRows(pods []corev1.Pod) (podList []table.Row) {
	for _, pod := range pods {
		podList = append(podList,
			table.Row{
				pod.Name,
				checkReadyContainers(
					pod.Status.ContainerStatuses),
				pod.Status.Reason,
				checkRestartedContainers(
					pod.Status.ContainerStatuses),
				pod.Spec.NodeName,
				DeltaTime(
					pod.GetCreationTimestamp().Time),
			},
		)
	}
	return podList
}

func FetchColumns(pods []corev1.Pod) (podAttributes []table.Column) {
	// readyContainersFieldName := pods[0].Status.ContainerStatuses[0].Ready

	return append(podAttributes,
		table.Column{Title: "Name", Width: podFieldWidth("Name", pods)},
		table.Column{Title: "Ready", Width: 10},
		table.Column{Title: "Status", Width: 20},
		table.Column{Title: "Restarts", Width: 10},
		table.Column{Title: "Node", Width: podFieldWidth("pods[0].Spec.NodeName", pods)},
		table.Column{Title: "Age", Width: podFieldWidth("Age", pods)},
	)
}

func podFieldWidth(fieldName string, pods []corev1.Pod) int {
	return lo.Reduce(pods, func(width int, pod corev1.Pod, _ int) int {
		logging.Log.Info(fieldName, " ", pods[0].Spec.NodeName)

		fieldValue, ok := funk.Get(pod, fieldName).(string)
		if !ok {
			return width
		}
		if len(fieldValue) > width {
			return len(fieldValue)
		}
		return width
	}, 0)
}

func DeltaTime(t time.Time) string {
	return time.Since(t).String()
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
