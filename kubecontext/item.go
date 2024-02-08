package kubecontext

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

type contextItem struct {
	api.Context
}

func (ci contextItem) Title() string       { return ci.Cluster }
func (ci contextItem) FilterValue() string { return "" }
func (ci contextItem) Description() string {
	var user = ""
	var namespace = ""

	user = userStyle.Render(fmt.Sprintf("User: %s ", ci.AuthInfo))
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

func fetchAllContexts() []list.Item {
	var kubeconfig string
	var fileContent []byte
	var config = map[string]interface{}{}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	fileContent, _ = os.ReadFile(kubeconfig)
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		panic(err)
	}
	return newContextList(config)
}

func convertMapToContext(contextList []interface{}) []api.Context {
	return lo.Map(contextList, func(c interface{}, _ int) api.Context {
		contextMap := c.(map[string]interface{})
		contextDetails := contextMap["context"].(map[string]interface{})
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
