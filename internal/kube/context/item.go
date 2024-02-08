package kubecontext

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
	"k8s.io/client-go/tools/clientcmd/api"
)

type contextItem struct {
	api.Context
}

func (ci contextItem) Title() string       { return "Cluster: " + ci.Cluster }
func (ci contextItem) FilterValue() string { return "" }
func (ci contextItem) Description() string {
	var namespace = ""

	user := userStyle.Render(fmt.Sprintf("User: %s ", ci.AuthInfo))
	if ci.Namespace != "" {
		namespace = namespaceStyle.Render(fmt.Sprintf("Namespace: %s", ci.Namespace))
	}
	return user + namespace
}

func newContextList(config map[string]interface{}) []list.Item {
	contextList := config["contexts"].([]interface{})
	return lo.Map(
		convertMapToContext(contextList),
		func(c api.Context, _ int) list.Item {
			return contextItem{
				Context: c,
			}
		})
}

func convertMapToContext(contextList []interface{}) []api.Context {
	return lo.Map(contextList, func(c interface{}, _ int) api.Context {
		logging.Log.Info(c)

		contextMap := c.(map[interface{}]interface{})
		contextDetails := contextMap["context"].(map[interface{}]interface{})

		context := api.NewContext()

		context.AuthInfo = contextDetails["user"].(string)
		context.Cluster = contextDetails["cluster"].(string)

		if namespace, ok := contextDetails["namespace"].(string); ok {
			context.Namespace = namespace
		} else {
			context.Namespace = ""
		}

		return *context
	})
}
