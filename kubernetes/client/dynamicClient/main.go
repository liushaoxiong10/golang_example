package main

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 加载配置信息
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/liushaoxiong/.kube/config")
	if err != nil {
		panic(err)
	}
	// 实例化dynamic client
	dynamicClient, err := dynamic.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	// 设置gvr
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	ctx := context.Background()
	// 获取指定资源，得到 unstructured.UnstructuredList 指针类型
	unstructuredObj, err := dynamicClient.Resource(gvr).Namespace(apiv1.NamespaceAll).List(ctx, metav1.ListOptions{Limit: 20})
	if err != nil {
		panic(err)
	}

	podList := &corev1.PodList{}
	// unstructured.UnstructuredList 转 PodList
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}
	for _, pod := range podList.Items {
		fmt.Printf("namespace: %v\tname: %v\t status: %v\n", pod.Namespace, pod.Name, pod.Status.Phase)

	}

}
