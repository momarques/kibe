package kube

import (
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

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
}

func (p PodDescription) TabNames() []string {
	return LookupStructFieldNames(reflect.TypeOf(p))
}

func (p PodOverview) TabContent() string {
	ips := lo.Map(p.IPs, func(item net.IP, _ int) string {
		return item.String()
	})

	fieldNames := LookupStructFieldNames(reflect.TypeOf(p))

	// fieldNames = uistyles.ColorizeDescriptionSectionKeys(fieldNames)

	logging.Log.Info(fieldNames)

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
	return t.Render()
}

type PodOverview struct {
	Name           string   `kibedescription:"Name"`
	Namespace      string   `kibedescription:"Namespace"`
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
		ServiceAccount: pod.Spec.ServiceAccountName,
		IP:             net.ParseIP(pod.Status.PodIP),
		IPs: lo.Map(pod.Status.PodIPs, func(item corev1.PodIP, _ int) net.IP {
			return net.ParseIP(item.IP)
		}),
		ControlledBy: pod.OwnerReferences[0].Kind + "/" + pod.OwnerReferences[0].Name,
		QoSClass:     string(pod.Status.QOSClass),
	}
}
