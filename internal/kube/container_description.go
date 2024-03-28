package kube

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

type PodContainers []corev1.Container

func printContainerPorts(ports []corev1.ContainerPort) string {
	formattedPorts := lo.Map(ports, func(item corev1.ContainerPort, _ int) string {
		portValues := lo.Compact([]string{item.Name, value(item.Protocol), item.HostIP, value(item.HostPort), value(item.ContainerPort)})
		return strings.Join(portValues, "::")
	})
	return strings.Join(formattedPorts, "\n")
}

func printEnvVarSource(env *corev1.EnvVarSource) string {
	format := func(s1, s2 string) string {
		return fmt.Sprintf("%s.%s", s1, s2)
	}

	switch {
	case env.ConfigMapKeyRef != nil:
		return format(env.ConfigMapKeyRef.Name, env.ConfigMapKeyRef.Key)
	case env.FieldRef != nil:
		return format(env.FieldRef.APIVersion, env.FieldRef.FieldPath)
	case env.ResourceFieldRef != nil:
		return format(env.ResourceFieldRef.ContainerName, env.ResourceFieldRef.Resource)
	case env.SecretKeyRef != nil:
		return format(env.SecretKeyRef.Name, env.SecretKeyRef.Key)
	}
}

func printContainerEnvs(envs []corev1.EnvVar) string {
	envStrings := lo.Map(envs, func(item corev1.EnvVar, _ int) string {
		if item.ValueFrom != nil {
			return fmt.Sprintf("%s: %s", item.Name, printEnvVarSource(item.ValueFrom))
		}

		basicFormat := fmt.Sprintf("%s=%s", item.Name, item.Value)
		return basicFormat
	})
	return strings.Join(envStrings, "\n")
}

func printContainerEnvFrom(envFrom []corev1.EnvFromSource) string {
	envs := lo.Map(envFrom, func(item corev1.EnvFromSource, _ int) string {
		var prefix string = item.Prefix
		var refName string

		switch {
		case item.ConfigMapRef != nil:
			refName = item.ConfigMapRef.Name
		case item.SecretRef != nil:
			refName = item.SecretRef.Name
		}
		if prefix != "" {
			refName = prefix + " " + refName
		}
		return refName
	})
	return strings.Join(envs, "\n")
}

func printContainerDetails(c corev1.Container) string {
	keys := []string{
		"Image",
		"ImagePullPolicy",
		"WorkingDir",
		"Commands",
		"Args",
		"Ports",
		"Env",
		"EnvFrom",
		"Resources",
		"VolumeMounts",
		"VolumeDevices",
		"LivenessProbe",
		"ReadinessProbe",
		"StartupProbe",
		"SecurityContext",
		"Stdin",
		"StdinOnce",
		"TTY",
		"TerminationMessagePath",
		"TerminationMessagePolicy",
		"Lifecycle",
		"ResizePolicy",
		"RestartPolicy",
	}
	content := []string{
		c.Image,
		value(c.ImagePullPolicy),
		c.WorkingDir,
		style.FormatCommand(c.Command),
		style.FormatCommand(c.Args),
		printContainerPorts(c.Ports),
		printContainerEnvs(c.Env),
		printContainerEnvFrom(c.EnvFrom),
		value(c.Resources),
		value(c.VolumeMounts),
		value(c.VolumeDevices),
		value(c.LivenessProbe),
		value(c.ReadinessProbe),
		value(c.StartupProbe),
		value(c.SecurityContext),
		value(c.Stdin),
		value(c.StdinOnce),
		value(c.TTY),
		value(c.TerminationMessagePath),
		value(c.TerminationMessagePolicy),
		value(c.Lifecycle),
		value(c.ResizePolicy),
		value(c.RestartPolicy)}

	return style.FormatTable(keys, content)
}

func (pc PodContainers) fetchContainersAsString() []map[string]string {
	return lo.Map(pc, func(item corev1.Container, _ int) map[string]string {
		return map[string]string{
			item.Name: printContainerDetails(item),
		}
	})
}

func (pc PodContainers) TabContent(page int) string {
	containers := pc.fetchContainersAsString()
	if pc == nil || len(containers) == 0 {
		return ""
	}
	containerDetails := lo.Entries(containers[page])[0]

	return lipgloss.JoinVertical(
		lipgloss.Left,
		style.CoreHeaderTitleStyle().
			Render(containerDetails.Key),
		containerDetails.Value,
	)
}
