package kube

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DescribePod(c *ClientReady, podID string) *corev1.Pod {
	pod, err := c.Client.
		CoreV1().
		Pods(string(c.NamespaceSelected)).
		Get(context.Background(), podID, v1.GetOptions{})
	if err != nil {
		logging.Log.Error(err)
	}
	return pod
}

// PodDescription provides information about the pod structured in Sections
//
// Those Sections are segmented in categories to enable a cleaner view of all the pod config
// Every Section has its own style
type PodDescription struct {
	Overview      PodOverview         `kibedescription:"Overview"`
	Status        PodStatus           `kibedescription:"Status"`
	Labels        ResourceLabels      `kibedescription:"Labels"`
	Annotations   ResourceAnnotations `kibedescription:"Annotations"`
	Volumes       PodVolumes          `kibedescription:"Volumes"`
	Containers    PodContainers       `kibedescription:"Containers"`
	NodeSelectors PodNodeSelector     `kibedescription:"Node Selectors"`
	Tolerations   PodTolerations      `kibedescription:"Tolerations"`
	Events        []string            `kibedescription:"Events"`
}

func (p Pod) Describe(c *ClientReady, podID string) ResourceDescription {
	pod := DescribePod(c, podID)

	logging.Log.Info("conditions -> ", newPodStatus(pod))
	return PodDescription{
		Overview:      newPodOverview(pod),
		Status:        newPodStatus(pod),
		Labels:        ResourceLabels(pod.Labels),
		Annotations:   ResourceAnnotations(pod.Annotations),
		Volumes:       PodVolumes(pod.Spec.Volumes),
		Containers:    PodContainers(pod.Spec.Containers),
		NodeSelectors: PodNodeSelector(pod.Spec.NodeSelector),
		Tolerations:   PodTolerations(pod.Spec.Tolerations),
	}
}

func (pd PodDescription) TabNames() []string {
	return LookupStructFieldNames(reflect.TypeOf(pd))
}

func (pd PodDescription) TabContent() []string {
	return []string{
		pd.Overview.TabContent(),
		pd.Status.TabContent(),
		pd.Labels.TabContent(),
		pd.Annotations.TabContent(),
		pd.Volumes.TabContent(),
		pd.Containers.TabContent(),
		pd.NodeSelectors.TabContent(),
		pd.Tolerations.TabContent(),
		"",
	}
}

// PodOverview provides basic information about the pod
//
// This object must return the whole content in a single formatted string
type PodOverview struct {
	Name           string   `kibedescription:"Name"`
	Namespace      string   `kibedescription:"Namespace"`
	NodeName       string   `kibedescription:"Node Name"`
	ServiceAccount string   `kibedescription:"Service Account"`
	IP             net.IP   `kibedescription:"IP"`
	IPs            []net.IP `kibedescription:"IPs"`
	ControlledBy   string   `kibedescription:"Controlled by"`
	QoSClass       string   `kibedescription:"QoS Class"`
}

func getPodOwner(pod *corev1.Pod) string {
	if len(pod.OwnerReferences) > 0 {
		return pod.OwnerReferences[0].Kind + "/" + pod.OwnerReferences[0].Name
	}
	return ""
}

func newPodOverview(pod *corev1.Pod) PodOverview {
	return PodOverview{
		Name:           pod.Name,
		Namespace:      pod.Namespace,
		NodeName:       pod.Spec.NodeName,
		ServiceAccount: pod.Spec.ServiceAccountName,
		IP:             net.ParseIP(pod.Status.PodIP),
		IPs: lo.Map(pod.Status.PodIPs,
			func(item corev1.PodIP, _ int) net.IP {
				return net.ParseIP(item.IP)
			}),
		ControlledBy: getPodOwner(pod),
		QoSClass:     string(pod.Status.QOSClass),
	}
}

func (po PodOverview) TabContent() string {
	ips := lo.Map(po.IPs,
		func(item net.IP, _ int) string {
			return item.String()
		})

	fieldNames := LookupStructFieldNames(reflect.TypeOf(po))

	t := table.New()
	t.Rows(
		[]string{fieldNames[0], po.Name},
		[]string{fieldNames[1], po.Namespace},
		[]string{fieldNames[2], po.NodeName},
		[]string{fieldNames[3], po.ServiceAccount},
		[]string{fieldNames[4], po.IP.String()},
		[]string{fieldNames[5], strings.Join(ips, ",")},
		[]string{fieldNames[6], po.ControlledBy},
		[]string{fieldNames[7], po.QoSClass},
	)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

// PodStatus provides historic status information from the pod
type PodStatus struct {
	Start      time.Time `kibedescription:"Started at"`
	Status     string    `kibedescription:"Status"`
	Conditions []string  `kibedescription:"Conditions"`
}

func newPodStatus(pod *corev1.Pod) PodStatus {
	return PodStatus{
		Start:      pod.CreationTimestamp.Time,
		Status:     string(pod.Status.Phase),
		Conditions: lo.Map(pod.Status.Conditions, podConditionToString),
	}
}

func podConditionToString(condition corev1.PodCondition, _ int) string {
	questionCondition := fmt.Sprintf("%s?", condition.Type)

	switch condition.Status {
	case corev1.ConditionTrue:
		return uistyles.OKStatusMessage.Render(questionCondition)
	case corev1.ConditionFalse:
		return uistyles.NOKStatusMessage.Render(questionCondition)
	case corev1.ConditionUnknown:
		return uistyles.WarnStatusMessage.Render(questionCondition)
	}
	return ""
}

func (ps PodStatus) TabContent() string {
	fieldNames := LookupStructFieldNames(reflect.TypeOf(ps))

	conditionsValue := strings.Join(ps.Conditions, " -> ")

	t := table.New()
	t.Rows(
		[]string{fieldNames[0], ps.Start.String()},
		[]string{fieldNames[1], ps.Status},
		[]string{fieldNames[2], conditionsValue},
	)
	t.StyleFunc(func(row, col int) lipgloss.Style {
		if col == 1 {
			return lipgloss.NewStyle().Width(len(conditionsValue))
		}
		return uistyles.ColorizeTabKey(row, col)
	})
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

type PodContainers []corev1.Container

func (pc PodContainers) TabContent() string {
	t := table.New()
	t.Rows(pc.podContainerToTableRows()...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func (pc PodContainers) podContainerToTableRows() [][]string {
	return lo.Map(pc,
		func(c corev1.Container, index int) []string {
			return []string{fmt.Sprintf("Container %d", index), c.Name}
		})
}

type PodNodeSelector map[string]string

func (pn PodNodeSelector) TabContent() string {
	t := table.New()

	t.Rows(mapToTableRows(pn)...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

type PodTolerations []corev1.Toleration

func (pt PodTolerations) podTolerationsToTableRows() [][]string {
	return lo.Map(pt,
		func(t corev1.Toleration, index int) []string {
			return []string{t.Key, prettyPrintTolerations(t)}
		})
}

func (pt PodTolerations) TabContent() string {
	t := table.New()

	t.Rows(pt.podTolerationsToTableRows()...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func prettyPrintTolerations(t corev1.Toleration) string {
	toleration := strings.Builder{}

	if !lo.IsEmpty(t.Value) {
		toleration.WriteString(fmt.Sprintf("%s ", t.Value))
	}
	if !lo.IsEmpty(t.Effect) {
		toleration.WriteString(fmt.Sprintf("%s ", t.Effect))
	}
	if !lo.IsEmpty(t.Operator) {
		toleration.WriteString(fmt.Sprintf("op=%s ", t.Operator))
	}
	if !lo.IsEmpty(t.TolerationSeconds) {
		toleration.WriteString(fmt.Sprintf("for %ds ", *t.TolerationSeconds))
	}

	return toleration.String()
}
