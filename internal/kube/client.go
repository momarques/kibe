package kube

import (
	"context"
	"fmt"
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

type TableResponse struct {
	Columns   []table.Column
	Rows      []table.Row
	Operation string

	FetchDuration time.Duration
	Err           error
}

type ClientReady struct {
	*kubernetes.Clientset

	ContextSelected
	NamespaceSelected
	ResourceSelected

	Ctx    context.Context
	Cancel context.CancelFunc
	TableResponse
}

func NewClientReady(ctx context.Context) ClientReady {
	client := ClientReady{}
	client.Ctx, client.Cancel = context.WithCancel(ctx)
	return client
}

func (c ClientReady) WithContext(ctx context.Context) ClientReady {
	c.Ctx, c.Cancel = context.WithCancel(ctx)
	return c
}

func (c ClientReady) WithClusterContext(clusterContext string) ClientReady {
	c.Clientset = NewKubeClient(clusterContext)
	c.ContextSelected = ContextSelected(clusterContext)
	return c
}

func (c ClientReady) WithNamespace(namespace string) ClientReady {
	if namespace == "" {
		namespace = "default"
	}
	c.NamespaceSelected = NamespaceSelected(namespace)
	return c
}

func (c ClientReady) WithResource(r Resource) ClientReady {
	c.ResourceSelected = r
	return c
}

func (c ClientReady) FetchTableView() TableResponse {
	logging.Log.Info("client ", c)
	var now = time.Now()

	resource, err := c.ResourceSelected.List(c)

	return TableResponse{
		Columns:       resource.Columns(),
		Rows:          resource.Rows(),
		FetchDuration: time.Since(now),
		Err:           err,
	}
}

func (c ClientReady) FetchTableViewAsync(responseCh chan TableResponse) {
	var now = time.Now()

	resource, err := c.ResourceSelected.List(c)

	responseCh <- TableResponse{
		Columns:       resource.Columns(),
		Rows:          resource.Rows(),
		FetchDuration: time.Since(now),
		Err:           err,
	}
}

func (c *ClientReady) LogOperation() string {
	return fmt.Sprintf("listing %ss", c.ResourceSelected.Kind())
}
