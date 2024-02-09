package kube

import (
	"os"
	"path/filepath"

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

func FetchKubeConfig() api.Config {
	fileContent, err := os.ReadFile(RetrieveKubeConfigFilePath())
	if err != nil {
		panic(err)
	}
	configBytes, err := clientcmd.NewClientConfigFromBytes(fileContent)
	if err != nil {
		panic(err)
	}
	config, err := configBytes.RawConfig()
	if err != nil {
		panic(err)
	}
	return config
}
