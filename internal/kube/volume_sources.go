package kube

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

func printAWSElasticBlockStore(v *corev1.AWSElasticBlockStoreVolumeSource) string {
	vs := fmt.Sprintf("  Type: \tAWSElasticBlockStoreVolumeSource\n"+
		"	VolumeID: \t%v\n"+
		"	FSType: \t%v\n"+
		"	Partition: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.VolumeID, v.FSType, v.Partition, v.ReadOnly)
	return vs
}

func printAzureDisk(v *corev1.AzureDiskVolumeSource) string {
	vs := fmt.Sprintf("Type: \tAzureDisk\n"+
		"	DiskName: \t%v\n"+
		"	DataDiskURI: \t%v\n"+
		"	CachingMode: \t%v\n"+
		"	Kind: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.DiskName, v.DataDiskURI, v.CachingMode, v.Kind, v.ReadOnly)
	return vs
}
func printCSI(v *corev1.CSIVolumeSource) string {
	vs := fmt.Sprintf("Type: \tCSI\n"+
		"	Driver: \t%v\n"+
		"	ReadOnly: \t%v\n"+
		"	VolumeAttributes: \t%v\n"+
		"	NodePublishSecretRef: \t%v\n"+
		"	FSType: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.Driver, v.ReadOnly, v.VolumeAttributes, v.NodePublishSecretRef, v.FSType, v.ReadOnly)
	return vs
}
func printCephFS(v *corev1.CephFSVolumeSource) string {
	vs := fmt.Sprintf("Type: \tCephFS\n"+
		"	Monitors: \t%v\n"+
		"	Path: \t%v\n"+
		"	User: \t%v\n"+
		"	SecretFile: \t%v\n"+
		"	SecretRef: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.Monitors, v.Path, v.User, v.SecretFile, v.SecretRef, v.ReadOnly)
	return vs
}
func printCinder(v *corev1.CinderVolumeSource) string {
	vs := fmt.Sprintf("Type: \tCinder\n"+
		"	VolumeID: \t%v\n"+
		"	FSType: \t%v\n"+
		"	ReadOnly: \t%v\n"+
		"	SecretRef: \t%v\n",
		v.VolumeID, v.FSType, v.ReadOnly, v.SecretRef)
	return vs
}

func printConfigMap(v *corev1.ConfigMapVolumeSource) string {
	items := lo.Map(v.Items, func(item corev1.KeyToPath, _ int) string {
		return fmt.Sprintf("%s:%s:%d", item.Key, item.Path, item.Mode)
	})

	vs := fmt.Sprintf("Type: \tConfigMap\n"+
		"	Name: \t%s\n"+
		"	Items: \t%s\n"+
		"	DefaultMode: \t%v\n"+
		"	Optional: \t%v\n",
		v.Name, strings.Join(items, " :: "), *v.DefaultMode, v.Optional)
	return vs
}

func printDownwardAPI(v *corev1.DownwardAPIVolumeSource) string {
	vs := fmt.Sprintf("Type: \tDownwardAPI\n"+
		"	Items: \t%v\n"+
		" 	DefaultMode: \t%v\n",
		v.Items, v.DefaultMode)
	return vs
}
func printEmptyDir(v *corev1.EmptyDirVolumeSource) string {
	vs := fmt.Sprintf("Type: \tEmptyDir\n"+
		"	Medium: \t%v\n"+
		"	SizeLimit: \t%v\n",
		v.Medium, v.SizeLimit)
	return vs
}
func printEphemeral(v *corev1.EphemeralVolumeSource) string {
	vs := fmt.Sprintf("Type: \tEphemeral\n"+
		"	VolumeClaimTemplate: \t%v\n",
		v.VolumeClaimTemplate)
	return vs
}
func printFC(v *corev1.FCVolumeSource) string {
	vs := fmt.Sprintf("Type: \tFC\n"+
		"	TargetWWNs: \t%v\n"+
		"	Lun: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.TargetWWNs, v.Lun, v.ReadOnly)
	return vs
}
func printFlexVolume(v *corev1.FlexVolumeSource) string {
	vs := fmt.Sprintf("Type: \tFlexVolume\n"+
		"	Driver: \t%v\n"+
		"	FSType: \t%v\n"+
		"	SecretRef: \t%v\n"+
		"	ReadOnly: \t%v\n"+
		"	Options: \t%v\n",
		v.Driver, v.FSType, v.SecretRef, v.ReadOnly, v.Options)
	return vs
}
func printFlocker(v *corev1.FlockerVolumeSource) string {
	vs := fmt.Sprintf("Type: \tFlocker\n"+
		"	DatasetName: \t%v\n",
		v.DatasetName)
	return vs
}
func printGCEPersistentDisk(v *corev1.GCEPersistentDiskVolumeSource) string {
	vs := fmt.Sprintf("Type: \tGCEPersistentDisk\n"+
		"	PDName: \t%v\n"+
		"	FSType: \t%v\n"+
		"	Partition: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.PDName, v.FSType, v.Partition, v.ReadOnly)
	return vs
}
func printGitRepo(v *corev1.GitRepoVolumeSource) string {
	vs := fmt.Sprintf("Type: \tGitRepo\n"+
		"	Repository: \t%v\n"+
		"	Revision: \t%v\n"+
		"	Directory: \t%v\n",
		v.Repository, v.Revision, v.Directory)
	return vs
}
func printGlusterfs(v *corev1.GlusterfsVolumeSource) string {
	vs := fmt.Sprintf("Type: \tGlusterfs\n"+
		"	EndpointsName: \t%v\n"+
		"	Path: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.EndpointsName, v.Path, v.ReadOnly)
	return vs
}
func printHostPath(v *corev1.HostPathVolumeSource) string {
	vs := fmt.Sprintf("  Type: \tHostPath\n"+
		"	Path: \t%v\n"+
		"	HostPathType: \t%v\n",
		v.Path, *v.Type)
	return vs
}
func printNFS(v *corev1.NFSVolumeSource) string {
	vs := fmt.Sprintf("Type: \tNFS\n"+
		"	Server: \t%v\n"+
		"	Path: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.Server, v.Path, v.ReadOnly)
	return vs
}
func printPersistentVolumeClaim(v *corev1.PersistentVolumeClaimVolumeSource) string {
	vs := fmt.Sprintf("Type: \tPersistentVolumeClaim\n"+
		"	ClaimName: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.ClaimName, v.ReadOnly)
	return vs
}
func printPhotonPersistentDisk(v *corev1.PhotonPersistentDiskVolumeSource) string {
	vs := fmt.Sprintf("Type: \tPhotonPersistentDisk\n"+
		"	PdID: \t%v\n"+
		"	FSType: \t%v\n",
		v.PdID, v.FSType)
	return vs
}
func printPortworxVolume(v *corev1.PortworxVolumeSource) string {
	vs := fmt.Sprintf("Type: \tPortworxVolume\n"+
		"	VolumeID: \t%v\n"+
		"	FSType: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.VolumeID, v.FSType, v.ReadOnly)
	return vs
}

func printProjected(v *corev1.ProjectedVolumeSource) string {
	sources := lo.Map(v.Sources, func(item corev1.VolumeProjection, _ int) string {
		if item.ConfigMap != nil {
			return fmt.Sprintf("    ConfigMapName: \t%s\n"+
				"        ConfigMapOptionalName: \t%v\n",
				item.ConfigMap.Name, item.ConfigMap.Optional)
		} else if item.DownwardAPI != nil {
			return "DownwardAPI: \ttrue\n"
		} else if item.Secret != nil {
			return fmt.Sprintf("    SecretName: \t%s\n"+
				"        SecretOptionalName: \t%v\n",
				item.Secret.Name, item.Secret.Optional)
		} else if item.ServiceAccountToken != nil {
			return fmt.Sprintf("    TokenExpirationSeconds: \t%d\n"+
				"        TokenPath: \t%v\n",
				*item.ServiceAccountToken.ExpirationSeconds, item.ServiceAccountToken.Path)
		} else if item.ClusterTrustBundle != nil {
			return fmt.Sprintf("    Cluster: \t%s\n", *item.ClusterTrustBundle.Name)
		}
		logging.Log.Error("Unknown projected source")
		return ""
	})

	vs := fmt.Sprintf("Type: \tProjected\n"+
		"	Sources: \t%v\n"+
		"	DefaultMode: \t%v\n",
		sources, *v.DefaultMode)

	return vs
}
func printQuobyte(v *corev1.QuobyteVolumeSource) string {
	vs := fmt.Sprintf("Type: \tQuobyte\n"+
		"	Registry: \t%v\n"+
		"	Volume: \t%v\n"+
		"	ReadOnly: \t%v\n"+
		"	User: \t%v\n"+
		"	Group: \t%v\n"+
		"	Tenant: \t%v\n",
		v.Registry, v.Volume, v.ReadOnly, v.User, v.Group, v.Tenant)
	return vs
}
func printRBD(v *corev1.RBDVolumeSource) string {
	vs := fmt.Sprintf("Type: \tRBD\n"+
		"	CephMonitors: \t%v\n"+
		"	RBDImage: \t%v\n"+
		"	FSType: \t%v\n"+
		"	RadosPool: \t%v\n"+
		"	RBDKeyring: \t%v\n"+
		"	RBDUser: \t%v\n"+
		"	Keyring: \t%v\n"+
		"	SecretRef: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.CephMonitors, v.RBDImage, v.FSType, v.RBDPool, v.Keyring, v.RadosUser, v.Keyring, v.SecretRef, v.ReadOnly)
	return vs
}
func printScaleIO(v *corev1.ScaleIOVolumeSource) string {
	vs := fmt.Sprintf("Type: \tScaleIO\n"+
		"	Gateway: \t%v\n"+
		"	System: \t%v\n"+
		"	SecretRef: \t%v\n"+
		"	SSLEnabled: \t%v\n"+
		"	ProtectionDomain: \t%v\n"+
		"	StoragePool: \t%v\n"+
		"	VolumeName: \t%v\n"+
		"	FSType: \t%v\n"+
		"	ReadOnly: \t%v\n",
		v.Gateway, v.System, v.SecretRef, v.SSLEnabled, v.ProtectionDomain, v.StoragePool, v.VolumeName, v.FSType, v.ReadOnly)
	return vs
}
func printSecret(v *corev1.SecretVolumeSource) string {
	vs := fmt.Sprintf("Type: \tSecret\n"+
		"	SecretName: \t%v\n"+
		"	Items: \t%v\n"+
		"	DefaultMode: \t%v\n"+
		"	Optional: \t%v\n",
		v.SecretName, v.Items, *v.DefaultMode, v.Optional)
	return vs
}
func printStorageOS(v *corev1.StorageOSVolumeSource) string {
	vs := fmt.Sprintf("Type: \tStorageOS\n"+
		"	VolumeName: \t%v\n"+
		"	VolumeNamespace: \t%v\n"+
		"	FSType: \t%v\n"+
		"	ReadOnly: \t%v\n"+
		"	SecretRef: \t%v\n"+
		"	LocalObjectReference: \t%v\n",
		v.VolumeName, v.VolumeNamespace, v.FSType, v.ReadOnly, v.SecretRef, v.SecretRef.Name)
	return vs
}
func printVsphereVolume(v *corev1.VsphereVirtualDiskVolumeSource) string {
	vs := fmt.Sprintf("Type: \tVsphereVolume\n"+
		"	VolumePath: \t%v\n"+
		"	FSType: \t%v\n"+
		"	StoragePolicyName: \t%v\n"+
		"	StoragePolicyID: \t%v\n",
		v.VolumePath, v.FSType, v.StoragePolicyName, v.StoragePolicyID)
	return vs
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
			item.Name: printVolumeSource(item.VolumeSource),
		}
	})
}

func (pv PodVolumes) TabContent() string {
	volumes := pv.fetchVolumeSourcesAsString()

	volumeDetails := lo.Entries(volumes[0])[0]

	renderer, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(100),
		glamour.WithStyles(glamour.DarkStyleConfig),
	)
	if err != nil {
		logging.Log.Error(err)
	}
	out, err := renderer.Render(
		fmt.Sprintf("# %s\n ```yaml\n%s``` \n", volumeDetails.Key, volumeDetails.Value))

	if err != nil {
		logging.Log.Error(err)
	}
	return out
}

type VolumeDetails struct {
	Name    string
	Source  string
	Details interface{}
}

// func (pv PodVolumes) volumeSourceSliceToMap() map[string]string {
// 	return lo.SliceToMap(pv, func(item corev1.Volume) (string, string) {
// 		marshaled, err := json.Marshal(item.VolumeSource)
// 		if err != nil {
// 			logging.Log.Error(err)
// 		}
// 		return item.Name, string(marshaled)
// 	})
// }

// func removeNilSources(item map[string]interface{}, _ int) VolumeDetails {
// 	volume := lo.Entries(item)[0]
// 	logging.Log.Info("key ->> ", volume.Key)
// 	logging.Log.Info("value ->> ", volume.Value)

// 	var withoutNilSources = map[string]interface{}{}

// 	for k, v := range item {
// 		logging.Log.Info(k, "  ", v)

// 		withoutNilSource, ok := v.(map[string]interface{})
// 		if ok {
// 			withoutNilSources[k] = withoutNilSource
// 		}
// 	}

// 	volumeSource := lo.Entries(withoutNilSources)[0]
// 	return VolumeDetails{
// 		Name:    volume.Key,
// 		Source:  volumeSource.Key,
// 		Details: volumeSource.Value,
// 	}
// }

// func (pv PodVolumes) extractVolumeDetails() []VolumeDetails {
// 	return lo.Map(pv.volumeSourceSliceToMap(), removeNilSources)
// }

// func (pv PodVolumes) podVolumesToTableRowMap() map[string]string {
// 	volumeDetails := pv.extractVolumeDetails()

// 	return lo.SliceToMap(volumeDetails,
// 		func(item VolumeDetails) (string, string) {
// 			return item.Name, fmt.Sprintf("%s :: %v", item.Source, item.Details)
// 		})
// }
