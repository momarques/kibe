package kube

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

var ResquestTimeout = 5 * time.Second

func RetrieveKubeConfigFilePath() string {
	if home := homedir.HomeDir(); home != "" {
		return filepath.Join(home, ".kube", "config")
	}
	return ""
}

func FetchKubeConfig() *api.Config {
	config, err := clientcmd.LoadFromFile(RetrieveKubeConfigFilePath())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return config
}

func NewKubeRestConfig(context string) *rest.Config {
	var overrides *clientcmd.ConfigOverrides

	if context != "" {
		overrides = &clientcmd.ConfigOverrides{
			CurrentContext: context,
		}
	}

	clientConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: RetrieveKubeConfigFilePath()},
		overrides,
	).ClientConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return clientConfig
}

func NewKubeClient(context string) *kubernetes.Clientset {
	clientConfig := NewKubeRestConfig(context)

	clientConfig.Timeout = ResquestTimeout
	client, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}

type TableResponse struct {
	Columns   []table.Column
	Rows      []table.Row
	Operation string

	FetchDuration time.Duration
	FetchErr      error
}
