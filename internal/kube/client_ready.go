package kube

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
)

type ClientReady struct {
	*kubernetes.Clientset

	ContextSelected
	NamespaceSelected
	ResourceSelected

	Ctx    context.Context
	Cancel context.CancelFunc
	Err    chan error
	TableResponse
}

func NewClientReady() ClientReady {
	return ClientReady{
		Err: make(chan error),
	}
}

func (c ClientReady) WithContext() ClientReady {
	c.Ctx, c.Cancel = context.WithCancel(context.Background())
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
	var now = time.Now()

	resource, err := c.ResourceSelected.List(c)

	return TableResponse{
		Columns:       resource.Columns(),
		Rows:          resource.Rows(),
		FetchDuration: time.Since(now),
		FetchErr:      err,
	}
}

func (c ClientReady) FetchTableViewAsync(responseCh chan TableResponse) {
	var now = time.Now()

	resource, err := c.ResourceSelected.List(c)

	responseCh <- TableResponse{
		Columns:       resource.Columns(),
		Rows:          resource.Rows(),
		FetchDuration: time.Since(now),
		FetchErr:      err,
	}
}

func (c *ClientReady) LogOperation() string {
	return fmt.Sprintf("%ss listed", c.ResourceSelected.Kind())
}
