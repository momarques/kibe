package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SupportedResources = []Resource{
	NewPodResource(),
	NewNamespaceResource(),
	NewServiceResource(),
}

type Resource interface{ Kind() string }

type ResourceSectionDescription interface {
	TabNames() []string
}

type ResourceSectionContent interface {
	TabContent() string
}

type SelectResource struct{ Resources []list.Item }

func NewSelectResource(c *ClientReady) func() tea.Msg {
	return func() tea.Msg {
		return SelectResource{
			Resources: ListAvailableResources(c)}
	}
}

type ResourceSelected struct{ R Resource }

type ResourceItem struct{ apiVersion, kind string }

func (r ResourceItem) Title() string       { return r.kind }
func (r ResourceItem) FilterValue() string { return r.kind }
func (r ResourceItem) Description() string {
	return fmt.Sprintf("API Version: %s", r.apiVersion)
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

func ListAvailableResources(c *ClientReady) []list.Item {
	apiList, err := c.Client.ServerPreferredResources()
	if err != nil {
		logging.Log.Error(err)
	}
	return newResourceList(apiList)
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
