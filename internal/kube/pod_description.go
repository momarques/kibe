package kube

import (
	"encoding/json"
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
)

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

func NewPodDescription(c *ClientReady, podID string) PodDescription {
	pod := DescribePod(c, podID)

	return PodDescription{
		Overview:      newPodOverview(pod),
		Status:        newPodStatus(pod),
		Labels:        ResourceLabels(pod.Labels),
		Annotations:   ResourceAnnotations(pod.Annotations),
		Volumes:       newPodVolumes(pod),
		Containers:    newPodContainers(pod),
		NodeSelectors: newPodNodeSelector(pod),
		Tolerations:   newPodTolerations(pod),
	}
}

func (pd PodDescription) TabNames() []string {
	return LookupStructFieldNames(reflect.TypeOf(pd))
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
	ControlledBy   string   `kibedescription:"Controlled By"`
	QoSClass       string   `kibedescription:"QoS Class"`
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
		ControlledBy: pod.OwnerReferences[0].Kind + "/" + pod.OwnerReferences[0].Name,
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
		Start:  pod.CreationTimestamp.Time,
		Status: string(pod.Status.Phase),
		Conditions: lo.Map(
			pod.Status.Conditions,
			podConditionToString),
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

	t := table.New()
	t.Rows(
		[]string{fieldNames[0], ps.Start.String()},
		[]string{fieldNames[1], ps.Status},
		[]string{fieldNames[2], strings.Join(ps.Conditions, "\n")},
	)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

type PodVolumes []corev1.Volume

func newPodVolumes(pod *corev1.Pod) PodVolumes {
	return PodVolumes(pod.Spec.Volumes)
}

func (pv PodVolumes) TabContent() string {
	t := table.New()
	t.Rows(mapToTableRows(
		pv.podVolumesToTableRows())...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func (pv PodVolumes) podVolumesToTableRows() map[string]string {
	return lo.SliceToMap(pv,
		func(item corev1.Volume) (string, string) {
			jsonString, err := item.VolumeSource.Marshal()
			if err != nil {
				logging.Log.Error(err)
			}
			var volumeSourceAsMap map[string]interface{}

			err = json.Unmarshal(jsonString, &volumeSourceAsMap)
			if err != nil {
				logging.Log.Error(err)
			}

			return item.Name, fmt.Sprintf("%v", volumeSourceAsMap)
		})
}

type PodContainers []corev1.Container

func newPodContainers(pod *corev1.Pod) PodContainers {
	return PodContainers(pod.Spec.Containers)
}

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

func newPodNodeSelector(pod *corev1.Pod) PodNodeSelector {
	return PodNodeSelector(pod.Spec.NodeSelector)
}

func (pn PodNodeSelector) TabContent() string {
	t := table.New()

	t.Rows(mapToTableRows(pn)...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

type PodTolerations []corev1.Toleration

func newPodTolerations(pod *corev1.Pod) PodTolerations {
	return PodTolerations(pod.Spec.Tolerations)
}

func (pt PodTolerations) podTolerationsToTableRows() [][]string {
	return lo.Map(pt,
		func(t corev1.Toleration, index int) []string {
			return []string{fmt.Sprintf("%s=%s:%s op=%s for %ds",
				t.Key,
				t.Value,
				t.Effect,
				t.Operator,
				t.TolerationSeconds)}
		})
}

func (pt PodTolerations) TabContent() string {
	t := table.New()

	t.Rows(pt.podTolerationsToTableRows()...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}
