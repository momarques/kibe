package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/momarques/kibe/internal/logging"
	modelstyles "github.com/momarques/kibe/internal/model/styles"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SupportedResources = []Resource{
	NewPodResource(),
	NewNamespaceResource(),
	NewServiceResource(),
}

type Resource interface{ Kind() string }

type ResourceSelected struct{ Resource }

type SelectResource struct{ apiVersion, kind string }

func (ri SelectResource) Title() string       { return ri.kind }
func (ri SelectResource) FilterValue() string { return ri.kind }
func (ri SelectResource) Description() string {
	return modelstyles.UserStyle.Render(fmt.Sprintf("API Version: %s", ri.apiVersion))
}

func newResourceList(apiList []*v1.APIResourceList) []list.Item {
	resourceList := []list.Item{}

	for _, v := range SupportedResources {
		resourceList = append(resourceList, SelectResource{
			kind:       v.Kind(),
			apiVersion: LookupAPIVersion(v.Kind(), apiList),
		})
	}
	return resourceList
}

func ListAvailableResources(c *ClientReady) ([]list.Item, error) {
	apiList, err := c.Client.ServerPreferredResources()
	if err != nil {

		logging.Log.Error(err)
	}
	return newResourceList(apiList), nil
}

func LookupAPIVersion(kind string, apiList []*v1.APIResourceList) string {
	for _, v := range apiList {
		for _, r := range v.APIResources {
			if r.Kind == kind {
				return v.GroupVersion
			}
		}
	}
	return ""
}
