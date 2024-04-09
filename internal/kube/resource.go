package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Resource interface {
	Describe(ClientReady) ResourceDescription
	ID() string
	Kind() string
	SetID(string) Resource

	List(ClientReady) (Resource, error)
	Columns() []table.Column
	Rows() []table.Row
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

var SupportedResources = []Resource{
	NewPodResource(),
	NewNamespaceResource(),
	NewServiceResource(),
}

func ListAvailableResources(c ClientReady) []list.Item {
	apiList, err := c.ServerPreferredResources()
	if err != nil {
		c.Err <- err
	}
	return lo.Map(SupportedResources,
		func(item Resource, _ int) list.Item {
			return ResourceItem{
				kind:       item.Kind(),
				apiVersion: LookupAPIVersion(item.Kind(), apiList),
			}
		})
}

type SelectResource struct{ Resources []list.Item }

func NewSelectResource(c ClientReady) func() tea.Msg {
	return func() tea.Msg {
		return SelectResource{
			Resources: ListAvailableResources(c)}
	}
}

type ResourceSelected Resource

type ResourceItem struct{ apiVersion, kind string }

func (r ResourceItem) Title() string       { return r.kind }
func (r ResourceItem) FilterValue() string { return r.kind }
func (r ResourceItem) Description() string {
	return fmt.Sprintf("API Version: %s", r.apiVersion)
}
