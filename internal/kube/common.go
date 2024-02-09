package kube

import (
	"path/filepath"

	"github.com/momarques/kibe/internal/logging"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

func RetrieveKubeConfigFilePath() string {
	if home := homedir.HomeDir(); home != "" {
		return filepath.Join(home, ".kube", "config")
	}
	return ""
}

func FetchKubeConfig() *api.Config {
	config, err := clientcmd.LoadFromFile(RetrieveKubeConfigFilePath())
	if err != nil {
		logging.Log.Error(err)
	}
	return config
}

func NewKubeClient(contextName string) *kubernetes.Clientset {
	var overrides *clientcmd.ConfigOverrides

	if contextName != "" {
		overrides = &clientcmd.ConfigOverrides{CurrentContext: contextName}
	}

	clientConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: RetrieveKubeConfigFilePath()},
		overrides,
	).ClientConfig()
	if err != nil {
		logging.Log.Error(err)
	}

	client, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		logging.Log.Error(err)
	}
	return client
}
