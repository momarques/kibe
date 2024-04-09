package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
	"k8s.io/client-go/tools/clientcmd/api"
)

func CurrentContext() string { return FetchKubeConfig().CurrentContext }

func ListContexts(config *api.Config) []list.Item {
	return lo.Map(lo.Values(config.Contexts),
		func(item *api.Context, _ int) list.Item {
			return ContextItem(*item)
		})
}

type SelectContext struct {
	Contexts       []list.Item
	CurrentContext string
}

func NewSelectContext() func() tea.Msg {
	config := FetchKubeConfig()
	return func() tea.Msg {
		return SelectContext{
			Contexts:       ListContexts(config),
			CurrentContext: config.CurrentContext,
		}
	}
}

type ContextSelected string

func (c ContextSelected) String() string { return string(c) }

type ContextItem api.Context

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
