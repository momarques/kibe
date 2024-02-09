package kubeapis

import (
	"encoding/json"
	"fmt"

	"k8s.io/client-go/tools/clientcmd/api"
)

// func Connect(c api.Config) {
// 	// create the clientset

// 	fmt.Printf("%+v\n", p) // {Name:Amit Kumar Age:30}

// 	clientConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
// 		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kube.RetrieveKubeConfig()},
// 		&clientcmd.ConfigOverrides{
// 			CurrentContext: context,
// 		}).ClientConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	clientset, err := kubernetes.NewForConfig(clientConfig)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func AAAA() {

	// config := make(map[string]interface{})

	// fileContent, err := os.ReadFile(kube.RetrieveKubeConfig())
	// if err != nil {
	// 	panic(err)
	// }

	// a, err := clientcmd.NewClientConfigFromBytes(fileContent)
	// if err != nil {
	// 	panic(err)
	// }
	// c, err := a.RawConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(c.Contexts)

	// clientConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
	// 	&clientcmd.ClientConfigLoadingRules{ExplicitPath: kube.RetrieveKubeConfig()},
	// 	&clientcmd.ConfigOverrides{
	// 		CurrentContext: "",
	// 	}).ClientConfig()
	// if err != nil {
	// 	panic(err)
	// }
}

func convertMapToStruct(m interface{}) {

	var p api.Config

	a, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(a, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

}

func convertMap(input interface{}) interface{} {
	switch in := input.(type) {
	case map[interface{}]interface{}:
		out := make(map[string]interface{})
		for key, value := range in {
			out[fmt.Sprint(key)] = convertMap(value)
		}
		return out
	case []interface{}:
		for i, v := range in {
			in[i] = convertMap(v)
		}
	}
	return input
}
