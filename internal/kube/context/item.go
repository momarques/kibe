package kubecontext

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/momarques/kibe/internal/kube"
	modelstyles "github.com/momarques/kibe/internal/model/styles"
	"k8s.io/client-go/tools/clientcmd/api"
)

type ContextItem struct{ api.Context }

func (ci ContextItem) Title() string       { return "Cluster: " + ci.Cluster }
func (ci ContextItem) FilterValue() string { return ci.Cluster }
func (ci ContextItem) Description() string {
	var namespace = ""

	user := modelstyles.UserStyle.Render(fmt.Sprintf("User: %s ", ci.AuthInfo))
	if ci.Namespace != "" {
		namespace = modelstyles.NamespaceStyle.Render(fmt.Sprintf("Namespace: %s", ci.Namespace))
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

func FetchItems() ([]list.Item, error) { return newContextList(kube.FetchKubeConfig()), nil }
