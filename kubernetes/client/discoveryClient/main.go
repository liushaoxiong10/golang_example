package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置信息
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/liushaoxiong/.kube/config")
	if err != nil {
		panic(err)
	}
	// 实例化discovery client
	discoverClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}
	// 返回支持的资源组，资源版本，资源信息
	_, APIResponse, err := discoverClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}

	for _, list := range APIResponse {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}
		for _, resource := range list.APIResources {
			fmt.Printf("name: %v\t group: %v\tversion: %v\n", resource.Name, gv.Group, gv.Version)
		}
	}

}
