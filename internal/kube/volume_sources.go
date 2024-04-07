package kube

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style"
	"github.com/momarques/kibe/internal/ui/style/theme"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

func value(v any) string {
	return fmt.Sprintf("%v", v)
}

func printAWSElasticBlockStore(v *corev1.AWSElasticBlockStoreVolumeSource) string {
	keys := []string{
		"Type", "VolumeID", "FSType",
		"Partition", "ReadOnly",
	}
	content := []string{
		"AWSElasticBlockStoreVolumeSource", v.VolumeID, v.FSType,
		value(v.Partition), value(v.ReadOnly)}

	return theme.FormatTable(keys, content)
}

func printAzureDisk(v *corev1.AzureDiskVolumeSource) string {
	keys := []string{
		"Type", "DiskName", "DataDiskURI",
		"CachingMode", "Kind", "ReadOnly",
	}
	content := []string{
		"AzureDisk", v.DiskName, v.DataDiskURI,
		value(v.CachingMode), value(v.Kind), value(v.ReadOnly),
	}

	return theme.FormatTable(keys, content)
}

func printCSI(v *corev1.CSIVolumeSource) string {
	keys := []string{
		"Type", "Driver", "ReadOnly",
		"VolumeAttributes", "NodePublishSecretRef", "FSType",
	}
	content := []string{
		"CSI", v.Driver, value(v.ReadOnly),
		value(v.VolumeAttributes), v.NodePublishSecretRef.Name, *v.FSType,
	}

	return theme.FormatTable(keys, content)
}

func printCephFS(v *corev1.CephFSVolumeSource) string {
	keys := []string{
		"Type", "Monitors", "Path",
		"User", "SecretFile", "SecretRef", "ReadOnly",
	}
	content := []string{
		"CephFS", strings.Join(v.Monitors, ":"), v.Path,
		v.User, v.SecretFile, v.SecretRef.Name, value(v.ReadOnly),
	}

	return theme.FormatTable(keys, content)
}

func printCinder(v *corev1.CinderVolumeSource) string {
	keys := []string{
		"Type", "VolumeID", "FSType",
		"ReadOnly", "SecretRef",
	}
	content := []string{
		"Cinder", v.VolumeID, v.FSType,
		value(v.ReadOnly), v.SecretRef.Name,
	}

	return theme.FormatTable(keys, content)
}

func formatKeyToPath(item corev1.KeyToPath, _ int) string {
	return fmt.Sprintf("%s:%s:%d", item.Key, item.Path, item.Mode)
}

func printConfigMap(v *corev1.ConfigMapVolumeSource) string {
	keys := []string{
		"Type", "Name", "Items",
		"DefaultMode", "Optional",
	}

	configMaps := lo.Map(v.Items, formatKeyToPath)
	content := []string{
		"ConfigMap", v.Name, strings.Join(configMaps, " | "),
		value(*v.DefaultMode), value(v.Optional),
	}
	return theme.FormatTable(keys, content)
}

func formatDownwardAPIVolumeSource(item corev1.DownwardAPIVolumeFile, _ int) string {
	return fmt.Sprintf("%s:%s:%d", item.FieldRef, item.Path, *item.Mode)
}

func printDownwardAPI(v *corev1.DownwardAPIVolumeSource) string {
	keys := []string{"Type", "Items", "DefaultMode"}

	volumeFiles := lo.Map(v.Items, formatDownwardAPIVolumeSource)
	content := []string{"DownwardAPI", strings.Join(volumeFiles, " | "), value(v.DefaultMode)}

	return theme.FormatTable(keys, content)
}

func printEmptyDir(v *corev1.EmptyDirVolumeSource) string {
	keys := []string{"Type", "Medium", "SizeLimit"}
	content := []string{"EmptyDir", value(v.Medium), v.SizeLimit.String()}

	return theme.FormatTable(keys, content)
}

func printEphemeral(v *corev1.EphemeralVolumeSource) string {
	keys := []string{
		"Type", "VolumeClaimTemplateName", "VolumeClaimTemplateNamespace",
	}
	content := []string{
		"Ephemeral", v.VolumeClaimTemplate.Name, v.VolumeClaimTemplate.Namespace,
	}

	return theme.FormatTable(keys, content)
}

func printFC(v *corev1.FCVolumeSource) string {
	keys := []string{"Type", "TargetWWNs", "Lun", "ReadOnly"}
	content := []string{
		"FC", strings.Join(v.TargetWWNs, " | "), value(v.Lun), value(v.ReadOnly),
	}

	return theme.FormatTable(keys, content)
}

func printFlexVolume(v *corev1.FlexVolumeSource) string {
	keys := []string{
		"Type", "Driver", "FSType",
		"SecretRef", "ReadOnly", "Options",
	}
	content := []string{
		"FlexVolume", v.Driver, v.FSType,
		v.SecretRef.Name, value(v.ReadOnly), value(v.Options),
	}

	return theme.FormatTable(keys, content)
}

func printFlocker(v *corev1.FlockerVolumeSource) string {
	keys := []string{"Type", "DatasetName"}
	content := []string{"Flocker", v.DatasetName}

	return theme.FormatTable(keys, content)
}

func printGCEPersistentDisk(v *corev1.GCEPersistentDiskVolumeSource) string {
	keys := []string{"Type", "PDName", "FSType", "Partition", "ReadOnly"}
	content := []string{
		"GCEPersistentDisk", v.PDName, v.FSType, value(v.Partition), value(v.ReadOnly),
	}

	return theme.FormatTable(keys, content)
}

func printGitRepo(v *corev1.GitRepoVolumeSource) string {
	keys := []string{"Type", "Repository", "Revision", "Directory"}
	content := []string{"GitRepo", v.Repository, v.Revision, v.Directory}

	return theme.FormatTable(keys, content)
}

func printGlusterfs(v *corev1.GlusterfsVolumeSource) string {
	keys := []string{"Type", "EndpointsName", "Path", "ReadOnly"}
	content := []string{"Glusterfs", v.EndpointsName, v.Path, value(v.ReadOnly)}

	return theme.FormatTable(keys, content)
}

func printHostPath(v *corev1.HostPathVolumeSource) string {
	keys := []string{"Type", "Path", "HostPathType"}
	content := []string{"HostPath", v.Path, value(*v.Type)}

	return theme.FormatTable(keys, content)
}

func printNFS(v *corev1.NFSVolumeSource) string {
	keys := []string{"Type", "Server", "Path", "ReadOnly"}
	content := []string{"NFS", v.Server, v.Path, value(v.ReadOnly)}

	return theme.FormatTable(keys, content)
}

func printPersistentVolumeClaim(v *corev1.PersistentVolumeClaimVolumeSource) string {
	keys := []string{"Type", "ClaimName", "ReadOnly"}
	content := []string{"PersistentVolumeClaim", v.ClaimName, value(v.ReadOnly)}

	return theme.FormatTable(keys, content)
}

func printPhotonPersistentDisk(v *corev1.PhotonPersistentDiskVolumeSource) string {
	keys := []string{"Type", "PdID", "FSType"}
	content := []string{"PhotonPersistentDisk", v.PdID, v.FSType}

	return theme.FormatTable(keys, content)
}

func printPortworxVolume(v *corev1.PortworxVolumeSource) string {
	keys := []string{"Type", "VolumeID", "FSType", "ReadOnly"}
	content := []string{"PortworxVolume", v.VolumeID, v.FSType, value(v.ReadOnly)}

	return theme.FormatTable(keys, content)
}

func extractVolumeProjection(item corev1.VolumeProjection, _ int) string {
	if item.ConfigMap != nil {
		subKeys := []string{"ConfigMapName", "ConfigMapOptionalName"}
		subValues := []string{item.ConfigMap.Name, value(item.ConfigMap.Optional)}
		return theme.FormatSubTable(subKeys, subValues)

	} else if item.DownwardAPI != nil {
		subKeys := []string{"DownwardAPI"}
		subValues := []string{"true"}
		return theme.FormatSubTable(subKeys, subValues)

	} else if item.Secret != nil {
		subKeys := []string{"DownwardAPI"}
		subValues := []string{"true"}
		return theme.FormatSubTable(subKeys, subValues)

	} else if item.ServiceAccountToken != nil {
		subKeys := []string{"TokenExpirationSeconds", "TokenPath"}
		subValues := []string{
			value(*item.ServiceAccountToken.ExpirationSeconds),
			item.ServiceAccountToken.Path,
		}
		return theme.FormatSubTable(subKeys, subValues)

	} else if item.ClusterTrustBundle != nil {
		subKeys := []string{"Cluster"}
		subValues := []string{*item.ClusterTrustBundle.Name}
		return theme.FormatSubTable(subKeys, subValues)
	}
	return ""
}

func printProjected(v *corev1.ProjectedVolumeSource) string {
	keys := []string{"Type", "Sources", "DefaultMode"}

	sources := lo.Map(v.Sources, extractVolumeProjection)
	content := []string{"Projected", strings.Join(sources, "\n"), value(*v.DefaultMode)}

	return theme.FormatTable(keys, content)
}

func printQuobyte(v *corev1.QuobyteVolumeSource) string {
	keys := []string{
		"Type", "Registry", "Volume",
		"ReadOnly", "User", "Group", "Tenant",
	}
	content := []string{
		"Quobyte", v.Registry, v.Volume,
		value(v.ReadOnly), v.User, v.Group, v.Tenant,
	}

	return theme.FormatTable(keys, content)
}

func printRBD(v *corev1.RBDVolumeSource) string {
	keys := []string{
		"Type", "CephMonitors", "RBDImage",
		"FSType", "RadosPool", "RBDKeyring",
		"RBDUser", "Keyring", "SecretRef", "ReadOnly",
	}
	content := []string{
		"RBD", strings.Join(v.CephMonitors, " | "), v.RBDImage,
		v.FSType, v.RBDPool, v.Keyring,
		v.RadosUser, v.Keyring, v.SecretRef.Name, value(v.ReadOnly),
	}

	return theme.FormatTable(keys, content)
}

func printScaleIO(v *corev1.ScaleIOVolumeSource) string {
	keys := []string{
		"Type", "Gateway", "System",
		"SecretRef", "SSLEnabled", "ProtectionDomain",
		"StoragePool", "VolumeName", "FSType", "ReadOnly",
	}
	content := []string{
		"ScaleIO", v.Gateway, v.System,
		v.SecretRef.Name, value(v.SSLEnabled), v.ProtectionDomain,
		v.StoragePool, v.VolumeName, v.FSType, value(v.ReadOnly),
	}

	return theme.FormatTable(keys, content)
}

func printSecret(v *corev1.SecretVolumeSource) string {
	keys := []string{"Type", "SecretName", "Items", "DefaultMode", "Optional"}
	secrets := lo.Map(v.Items, formatKeyToPath)
	content := []string{
		"Secret", v.SecretName, strings.Join(secrets, " | "), value(*v.DefaultMode), value(v.Optional),
	}

	return theme.FormatTable(keys, content)
}

func printStorageOS(v *corev1.StorageOSVolumeSource) string {
	keys := []string{
		"Type", "VolumeName", "VolumeNamespace",
		"FSType", "ReadOnly", "SecretRef", "LocalObjectReference",
	}
	content := []string{
		"StorageOS", v.VolumeName, v.VolumeNamespace,
		v.FSType, value(v.ReadOnly), v.SecretRef.Name, v.SecretRef.Name,
	}

	return theme.FormatTable(keys, content)
}

func printVsphereVolume(v *corev1.VsphereVirtualDiskVolumeSource) string {
	keys := []string{
		"Type", "VolumePath", "FSType",
		"StoragePolicyName", "StoragePolicyID",
	}
	content := []string{
		"VsphereVolume", v.VolumePath, v.FSType, v.StoragePolicyName, v.StoragePolicyID,
	}

	return theme.FormatTable(keys, content)
}

func printVolumeSource(v corev1.VolumeSource) string {
	switch {
	case v.AWSElasticBlockStore != nil:
		return printAWSElasticBlockStore(v.AWSElasticBlockStore)
	case v.AzureDisk != nil:
		return printAzureDisk(v.AzureDisk)
	case v.CSI != nil:
		return printCSI(v.CSI)
	case v.CephFS != nil:
		return printCephFS(v.CephFS)
	case v.Cinder != nil:
		return printCinder(v.Cinder)
	case v.ConfigMap != nil:
		return printConfigMap(v.ConfigMap)
	case v.DownwardAPI != nil:
		return printDownwardAPI(v.DownwardAPI)
	case v.EmptyDir != nil:
		return printEmptyDir(v.EmptyDir)
	case v.Ephemeral != nil:
		return printEphemeral(v.Ephemeral)
	case v.FC != nil:
		return printFC(v.FC)
	case v.FlexVolume != nil:
		return printFlexVolume(v.FlexVolume)
	case v.Flocker != nil:
		return printFlocker(v.Flocker)
	case v.GCEPersistentDisk != nil:
		return printGCEPersistentDisk(v.GCEPersistentDisk)
	case v.GitRepo != nil:
		return printGitRepo(v.GitRepo)
	case v.Glusterfs != nil:
		return printGlusterfs(v.Glusterfs)
	case v.HostPath != nil:
		return printHostPath(v.HostPath)
	case v.NFS != nil:
		return printNFS(v.NFS)
	case v.PersistentVolumeClaim != nil:
		return printPersistentVolumeClaim(v.PersistentVolumeClaim)
	case v.PhotonPersistentDisk != nil:
		return printPhotonPersistentDisk(v.PhotonPersistentDisk)
	case v.PortworxVolume != nil:
		return printPortworxVolume(v.PortworxVolume)
	case v.Projected != nil:
		return printProjected(v.Projected)
	case v.Quobyte != nil:
		return printQuobyte(v.Quobyte)
	case v.RBD != nil:
		return printRBD(v.RBD)
	case v.ScaleIO != nil:
		return printScaleIO(v.ScaleIO)
	case v.Secret != nil:
		return printSecret(v.Secret)
	case v.StorageOS != nil:
		return printStorageOS(v.StorageOS)
	case v.VsphereVolume != nil:
		return printVsphereVolume(v.VsphereVolume)
	}
	return ""
}

type PodVolumes []corev1.Volume

func (pv PodVolumes) fetchVolumeSourcesAsString() []map[string]string {
	return lo.Map(pv, func(item corev1.Volume, _ int) map[string]string {
		return map[string]string{
			item.Name: lipgloss.NewStyle().
				Render(printVolumeSource(item.VolumeSource)),
		}
	})
}

func (pv PodVolumes) TabContent(page int) string {
	volumes := pv.fetchVolumeSourcesAsString()
	if pv == nil || len(volumes) == 0 {
		return ""
	}
	volumeDetails := lo.Entries(volumes[page])[0]

	return lipgloss.JoinVertical(
		lipgloss.Left,
		style.CoreHeaderTitleStyle().
			Render(volumeDetails.Key),
		volumeDetails.Value,
	)
}
