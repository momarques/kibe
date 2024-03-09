package kube

import (
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/momarques/kibe/internal/logging"
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
		logging.Log.Error(err)
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
		logging.Log.Error(err)
	}
	return clientConfig
}

func NewKubeClient(context string) *kubernetes.Clientset {
	clientConfig := NewKubeRestConfig(context)

	clientConfig.Timeout = ResquestTimeout
	client, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		logging.Log.Error(err)
	}
	return client
}

type RequestState chan error

type ClientReady struct {
	Client *kubernetes.Clientset

	*ContextSelected
	*ResourceSelected
}

func NewClientReady(context string) *ClientReady {
	return &ClientReady{
		Client: NewKubeClient(context),
		ContextSelected: &ContextSelected{
			C: context,
		},
	}
}

func (c *ClientReady) WithNamespace(namespace string) *ClientReady {
	if namespace == "" {
		namespace = "default"
	}
	c.Namespace = &NamespaceSelected{NS: namespace}
	return c
}

func (c *ClientReady) WithResource(r Resource) *ClientReady {
	c.ResourceSelected = &ResourceSelected{r}
	return c
}

func (c *ClientReady) FetchTableView() ([]table.Column, []table.Row) {
	switch c.ResourceSelected.R.(type) {
	case *Pod:
		pods := ListPods(c)
		podColumns := ListPodColumns(pods)

		return podColumns, RetrievePodListAsTableRows(pods)
	case *Namespace:
		ns := ListNamespaces(c)
		nsColumns := ListNamespaceColumns(ns)

		return nsColumns, RetrieveNamespaceListAsTableRows(ns)
	case *Service:
		svc := ListServices(c)
		svcColumns := ListServiceColumns(svc)

		return svcColumns, RetrieveServiceListAsTableRows(svc)
	}
	return nil, nil
}
