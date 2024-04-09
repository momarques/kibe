package kube

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/momarques/kibe/internal/ui/style"
	"github.com/momarques/kibe/internal/ui/style/theme"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DescribePod(c ClientReady) *corev1.Pod {
	pod, err := c.
		CoreV1().
		Pods(c.NamespaceSelected.String()).
		Get(c.Ctx, c.ID(), v1.GetOptions{})
	if err != nil {
		c.Err <- err
	}
	return pod
}

// PodDescription provides information about the pod structured in Sections
//
// Those Sections are segmented in categories to enable a cleaner view of all the pod config
// Every Section has its own style
type PodDescription struct {
	Overview       PodOverview         `kibedescription:"Overview"`
	Status         PodStatus           `kibedescription:"Status"`
	Labels         ResourceLabels      `kibedescription:"Labels"`
	Annotations    ResourceAnnotations `kibedescription:"Annotations"`
	Volumes        PodVolumes          `kibedescription:"Volumes"`
	Containers     PodContainers       `kibedescription:"Containers"`
	NodeScheduling PodNodeScheduling   `kibedescription:"Node Scheduling"`
	Events         []string            `kibedescription:"Events"`
}

func (p Pod) Describe(c ClientReady) ResourceDescription {
	pod := DescribePod(c)

	return PodDescription{
		Overview:       newPodOverview(pod),
		Status:         newPodStatus(pod),
		Labels:         ResourceLabels(pod.Labels),
		Annotations:    ResourceAnnotations(pod.Annotations),
		Volumes:        PodVolumes(pod.Spec.Volumes),
		Containers:     PodContainers(pod.Spec.Containers),
		NodeScheduling: newPodNodeScheduling(pod),
	}
}

func (pd PodDescription) TabNames() []string {
	return LookupStructFieldNames(pd)
}

func (pd PodDescription) TabContent() []string {
	return []string{
		pd.Overview.TabContent(),
		pd.Status.TabContent(),
		pd.Labels.TabContent(),
		pd.Annotations.TabContent(),
		pd.Volumes.TabContent(0),
		pd.Containers.TabContent(0),
		pd.NodeScheduling.TabContent(),
		"",
	}
}

// PodOverview provides basic information about the pod
//
// This object must return the whole content in a single formatted string
type PodOverview struct {
	Name           string   `kibedescription:"Name"`
	Namespace      string   `kibedescription:"Namespace"`
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

func parseIP(item corev1.PodIP, _ int) net.IP {
	return net.ParseIP(item.IP)
}

func newPodOverview(pod *corev1.Pod) PodOverview {
	return PodOverview{
		Name:           pod.Name,
		Namespace:      pod.Namespace,
		ServiceAccount: pod.Spec.ServiceAccountName,
		IP:             net.ParseIP(pod.Status.PodIP),
		IPs:            lo.Map(pod.Status.PodIPs, parseIP),
		ControlledBy:   getPodOwner(pod),
		QoSClass:       string(pod.Status.QOSClass),
	}
}

func (po PodOverview) TabContent() string {
	keys := LookupStructFieldNames(po)

	ips := lo.Map(po.IPs, func(item net.IP, _ int) string { return item.String() })
	content := []string{
		po.Name, po.Namespace, po.ServiceAccount,
		po.IP.String(), strings.Join(ips, ","), po.ControlledBy, po.QoSClass,
	}
	return theme.FormatTable(keys, content)
}

// PodStatus provides historic status information from the pod
type PodStatus struct {
	Start      time.Time `kibedescription:"Started at"`
	Status     string    `kibedescription:"Status"`
	Conditions []string  `kibedescription:"Conditions"`
}

func podConditionToString(condition corev1.PodCondition, _ int) string {
	questionCondition := fmt.Sprintf("%s?", condition.Type)

	switch condition.Status {
	case corev1.ConditionTrue:
		return style.OKStatusMessage().Render(questionCondition)
	case corev1.ConditionFalse:
		return style.NOKStatusMessage().Render(questionCondition)
	case corev1.ConditionUnknown:
		return style.WarnStatusMessage().Render(questionCondition)
	}
	return ""
}

func newPodStatus(pod *corev1.Pod) PodStatus {
	return PodStatus{
		Start:      pod.CreationTimestamp.Time,
		Status:     string(pod.Status.Phase),
		Conditions: lo.Map(pod.Status.Conditions, podConditionToString),
	}
}

func (ps PodStatus) TabContent() string {
	keys := LookupStructFieldNames(ps)

	conditionsValue := strings.Join(ps.Conditions, "\n")
	content := []string{ps.Start.String(), ps.Status, conditionsValue}

	return theme.FormatTable(keys, content)
}

type PodNodeSelector map[string]string

func (pn PodNodeSelector) TabContent() string {
	keys := lo.Keys(pn)
	content := lo.Values(pn)

	return theme.FormatSubTable(keys, content)
}

type PodTolerations []corev1.Toleration

func prettyPrintTolerations(t corev1.Toleration, _ int) string {
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

func (pt PodTolerations) TabContent() string {
	keys := lo.Map(pt,
		func(t corev1.Toleration, index int) string {
			return t.Key
		})
	content := lo.Map(pt, prettyPrintTolerations)

	return theme.FormatSubTable(keys, content)
}

type PodNodeScheduling struct {
	NodeName      string          `kibedescription:"Node Name"`
	NodeSelectors PodNodeSelector `kibedescription:"Node Selectors"`
	Tolerations   PodTolerations  `kibedescription:"Tolerations"`
}

func newPodNodeScheduling(pod *corev1.Pod) PodNodeScheduling {
	return PodNodeScheduling{
		NodeName:      pod.Spec.NodeName,
		NodeSelectors: PodNodeSelector(pod.Spec.NodeSelector),
		Tolerations:   PodTolerations(pod.Spec.Tolerations),
	}
}

func (pn PodNodeScheduling) TabContent() string {
	keys := []string{"Node name", "Node selectors", "Tolerations"}
	content := []string{
		pn.NodeName, pn.NodeSelectors.TabContent(), pn.Tolerations.TabContent(),
	}

	return theme.FormatTable(keys, content)
}

func (pd PodDescription) SubContent(subContentIndex int) []string {
	t := reflect.TypeFor[PodDescription]()
	field := t.Field(subContentIndex)

	switch field.Name {
	case "Volumes":
		return lo.Map(pd.Volumes,
			func(item corev1.Volume, index int) string {
				return pd.Volumes.TabContent(index)
			})
	case "Containers":
		return lo.Map(pd.Containers,
			func(item corev1.Container, index int) string {
				return pd.Containers.TabContent(index)
			})
	}
	return []string{}
}
