package kube

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

type PodDescription struct {
	Overview PodOverview `kibetab:"Overview"`
	Status   struct {
		Start      time.Time
		Status     string
		Conditions []map[string]interface{}
	} `kibetab:"Status"`
	LabelsAndAnnotations struct {
		Labels      map[string]interface{}
		Annotations map[string]interface{}
	} `kibetab:"Labels and Annotations"`
	Mounts struct {
		Volumes []map[string]interface{}
	} `kibetab:"Mounts"`
	Containers []struct{} `kibetab:"Containers"`
	Scheduling struct {
		Node         string
		NodeSelector map[string]interface{}
		Tolerations  map[string]interface{}
		NodeAffinity map[string]interface{}
		PodAffinity  map[string]interface{}
	} `kibetab:"Scheduling"`
}

func (p PodDescription) TabNames() []string {
	return LookupStructFieldNames(reflect.TypeOf(p))
}

func (p PodOverview) TabContent() string {
	ips := lo.Map(p.IPs, func(item net.IP, _ int) string {
		return item.String()
	})

	fieldNames := LookupStructFieldNames(reflect.TypeOf(p))

	fieldNames = uistyles.ColorizeDescriptionSectionKeys(fieldNames)

	return strings.Join([]string{
		fmt.Sprintf("%s=%s", fieldNames[0], p.Name),
		fmt.Sprintf("%s=%s", fieldNames[1], p.Namespace),
		fmt.Sprintf("%s=%s", fieldNames[2], p.ServiceAccount),
		fmt.Sprintf("%s=%s", fieldNames[3], p.IP.String()),
		fmt.Sprintf("%s=%s", fieldNames[4], strings.Join(ips, ",")),
		fmt.Sprintf("%s=%s", fieldNames[5], p.ControlledBy),
		fmt.Sprintf("%s=%s", fieldNames[6], p.QoSClass)}, "\n")
}

type PodOverview struct {
	Name           string
	Namespace      string
	ServiceAccount string
	IP             net.IP
	IPs            []net.IP
	ControlledBy   string
	QoSClass       string
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
