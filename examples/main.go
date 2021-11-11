package main

import (
	"fmt"
	"log"

	"github.com/wlynch/k8stoken"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, nil)
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	token, err := k8stoken.Token(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
}
