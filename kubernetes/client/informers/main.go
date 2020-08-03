package main

import (
	"log"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/liushaoxiong/.kube/config")
	if err != nil {
		panic(err)
	}
	// 创建 chientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	// 实例化 SharedInformer
	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	informer := sharedInformers.Core().V1().Pods().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		// 创建资源时触发
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New Pod add to store: %s", mObj.GetName())
		},
		// 更新资源时触发
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s pod updated to %s", oObj.GetName(), nObj.GetName())
		},
		// 删除资源时触发
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("pod delete from store: %s", mObj.GetName())
		},
	})
	informer.Run(stopCh)
}
