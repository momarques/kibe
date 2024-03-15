package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"k8s.io/client-go/tools/clientcmd/api"
)

type SelectContext struct {
	Contexts       []list.Item
	CurrentContext string
}

func NewSelectContext() func() tea.Msg {
	return func() tea.Msg {
		return SelectContext{
			Contexts:       ListContexts(),
			CurrentContext: CurrentContext(),
		}
	}
}

type ContextSelected string
type ContextItem struct{ api.Context }

func (c ContextItem) Title() string       { return "Cluster: " + c.Cluster }
func (c ContextItem) FilterValue() string { return c.Cluster }
func (c ContextItem) Description() string {
	var namespace = ""

	user := fmt.Sprintf("User: %s ", c.AuthInfo)
	if c.Namespace != "" {
		namespace = fmt.Sprintf(":::::: Namespace: %s", c.Namespace)
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
func CurrentContext() string    { return FetchKubeConfig().CurrentContext }
