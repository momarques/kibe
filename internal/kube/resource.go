package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	modelstyles "github.com/momarques/kibe/internal/model/styles"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SupportedResources = []Resource{
	NewPodResource(),
	NewNamespaceResource(),
	NewServiceResource(),
}

type Resource interface{ Kind() string }

type ResourceItem struct{ apiVersion, kind string }

func (ri ResourceItem) Title() string       { return ri.kind }
func (ri ResourceItem) FilterValue() string { return ri.kind }
func (ri ResourceItem) Description() string {
	return modelstyles.UserStyle.Render(fmt.Sprintf("API Version: %s", ri.apiVersion))
}

func newResourceList(apiList []*v1.APIResourceList) []list.Item {
	resourceList := []list.Item{}

	for _, v := range SupportedResources {
		resourceList = append(resourceList, ResourceItem{
			kind:       v.Kind(),
			apiVersion: LookupAPIVersion(v.Kind(), apiList),
		})
	}
	return resourceList
}

func FetchListItems(a []*v1.APIResourceList) ([]list.Item, error) { return newResourceList(a), nil }
