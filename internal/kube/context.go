package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"k8s.io/client-go/tools/clientcmd/api"
)

type SelectContext struct{ Contexts []list.Item }

func NewSelectContext() func() tea.Msg {
	return func() tea.Msg {
		return SelectContext{
			Contexts: ListContexts(),
		}
	}
}

type ContextSelected struct {
	C         string
	Namespace *NamespaceSelected
}

type ContextItem struct{ api.Context }

func (c ContextItem) Title() string       { return "Cluster: " + c.Cluster }
func (c ContextItem) FilterValue() string { return c.Cluster }
func (c ContextItem) Description() string {
	var namespace = ""

	user := uistyles.UserStyle.Render(fmt.Sprintf("User: %s ", c.AuthInfo))
	if c.Namespace != "" {
		namespace = uistyles.NamespaceStyle.Render(fmt.Sprintf("Namespace: %s", c.Namespace))
	}
	return user + namespace
}

func newContextList(config *api.Config) []list.Item {
	contextList := []list.Item{}

	for _, v := range config.Contexts {
		contextList = append(contextList, ContextItem{
			Context: *v,
		})
	}
	return contextList
}

func ListContexts() []list.Item { return newContextList(FetchKubeConfig()) }
