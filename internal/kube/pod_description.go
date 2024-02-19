package kube

import (
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

// PodDescription provides information about the pod structured in Sections
//
// Those Sections are segmented in categories to enable a cleaner view of all the pod config
// Every Section has its own style
type PodDescription struct {
	Overview PodOverview `kibedescription:"Overview"`
	Status   struct {
		Start      time.Time
		Status     string
		Conditions []map[string]interface{}
	} `kibedescription:"Status"`
	LabelsAndAnnotations struct {
		Labels      map[string]interface{}
		Annotations map[string]interface{}
	} `kibedescription:"Labels and Annotations"`
	Mounts struct {
		Volumes []map[string]interface{}
	} `kibedescription:"Mounts"`
	Containers []struct{} `kibedescription:"Containers"`
	Scheduling struct {
		Node         string
		NodeSelector map[string]interface{}
		Tolerations  map[string]interface{}
		NodeAffinity map[string]interface{}
		PodAffinity  map[string]interface{}
	} `kibedescription:"Scheduling"`
	Events []string `kibedescription:"Scheduling"`
}

func NewPodDescription(c *ClientReady, podID string) PodDescription {
	pod := DescribePod(c, podID)

	return PodDescription{
		Overview: newPodOverview(pod),
	}
}

func (p PodDescription) TabNames() []string {
	return LookupStructFieldNames(reflect.TypeOf(p))
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
	ControlledBy   string   `kibedescription:"Controlled By"`
	QoSClass       string   `kibedescription:"QoS Class"`
}

func (p PodOverview) TabContent() string {
	ips := lo.Map(p.IPs, func(item net.IP, _ int) string {
		return item.String()
	})

	fieldNames := LookupStructFieldNames(reflect.TypeOf(p))

	t := table.New()
	t.Rows(
		[]string{fieldNames[0], p.Name},
		[]string{fieldNames[1], p.Namespace},
		[]string{fieldNames[2], p.ServiceAccount},
		[]string{fieldNames[3], p.IP.String()},
		[]string{fieldNames[4], strings.Join(ips, ",")},
		[]string{fieldNames[5], p.ControlledBy},
		[]string{fieldNames[6], p.QoSClass},
	)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func newPodOverview(pod *corev1.Pod) PodOverview {
	return PodOverview{
		Name:           pod.Name,
		Namespace:      pod.Namespace,
		ServiceAccount: pod.Spec.ServiceAccountName,
		IP:             net.ParseIP(pod.Status.PodIP),
		IPs: lo.Map(pod.Status.PodIPs, func(item corev1.PodIP, _ int) net.IP {
			return net.ParseIP(item.IP)
		}),
		ControlledBy: pod.OwnerReferences[0].Kind + "/" + pod.OwnerReferences[0].Name,
		QoSClass:     string(pod.Status.QOSClass),
	}
}
