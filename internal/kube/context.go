package kube

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	modelstyles "github.com/momarques/kibe/internal/model/styles"
	"k8s.io/client-go/tools/clientcmd/api"
)

type ContextSelected struct {
	C  string
	NS NamespaceSelected
}

type SelectContext struct{ api.Context }

func (s SelectContext) Title() string       { return "Cluster: " + s.Cluster }
func (s SelectContext) FilterValue() string { return s.Cluster }
func (s SelectContext) Description() string {
	var namespace = ""

	user := modelstyles.UserStyle.Render(fmt.Sprintf("User: %s ", s.AuthInfo))
	if s.Namespace != "" {
		namespace = modelstyles.NamespaceStyle.Render(fmt.Sprintf("Namespace: %s", s.Namespace))
	}
	return user + namespace
}

func newContextList(config *api.Config) []list.Item {
	contextList := []list.Item{}

	for _, v := range config.Contexts {
		contextList = append(contextList, SelectContext{
			Context: *v,
		})
	}
	return contextList
}

func ListContexts() ([]list.Item, error) { return newContextList(FetchKubeConfig()), nil }
