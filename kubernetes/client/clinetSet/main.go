package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/liushaoxiong/.kube/config")
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	podClient := clientSet.CoreV1().Pods(corev1.NamespaceAll)

	ctx := context.Background()
	list, err := podClient.List(ctx, metav1.ListOptions{Limit: 200})
	if err != nil {
		panic(err)
	}
	for _, pod := range list.Items {
		fmt.Printf("namespace: %v\tname: %v\t status: %v\n", pod.Namespace, pod.Name, pod.Status.Phase)

	}

}
