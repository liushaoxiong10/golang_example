package main

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	matev1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/deprecated/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/liushaoxiong/.kube/config")
	if err != nil {
		panic(err)
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := &corev1.PodList{}
	err = restClient.Get().Namespace("").Resource("pods").VersionedParams(&matev1.ListOptions{Limit: 200}, scheme.ParameterCodec).Do(ctx).Into(result)
	if err != nil {
		panic(err)
	}
	for _, pod := range result.Items {
		fmt.Printf("namespace: %v\tname: %v\t status: %v\n", pod.Namespace, pod.Name, pod.Status.Phase)

	}

}
