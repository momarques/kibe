package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
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

// ResourceLabels provides a map of labels from the resource
type ResourceLabels map[string]string

func (rl ResourceLabels) TabContent() string {
	t := table.New()
	t.Rows(mapToTableRows(rl)...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

// ResourceAnnotations provides a map of annotations from the resource
type ResourceAnnotations map[string]string

func (ra ResourceAnnotations) TabContent() string {
	t := table.New()
	t.Rows([]string{})
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func mapToTableRows(m map[string]string) [][]string {
	return lo.MapToSlice(m, func(k, v string) []string {
		return []string{k, v}
	})
}
