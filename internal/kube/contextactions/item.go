package contextactions

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/momarques/kibe/internal/kube"
	"k8s.io/client-go/tools/clientcmd/api"
)

type contextItem struct {
	api.Context
}

func (ci contextItem) Title() string       { return "Cluster: " + ci.Cluster }
func (ci contextItem) FilterValue() string { return ci.Cluster }
func (ci contextItem) Description() string {
	var namespace = ""

	user := userStyle.Render(fmt.Sprintf("User: %s ", ci.AuthInfo))
	if ci.Namespace != "" {
		namespace = namespaceStyle.Render(fmt.Sprintf("Namespace: %s", ci.Namespace))
	}
	return user + namespace
}

func newContextList(config *api.Config) []list.Item {
	contextList := []list.Item{}

	for _, v := range config.Contexts {
		contextList = append(contextList, contextItem{
			Context: *v,
		})
	}
	return contextList
}

func (ca *contextActions) FetchListItems() ([]list.Item, error) {
	return newContextList(kube.FetchKubeConfig()), nil
}
