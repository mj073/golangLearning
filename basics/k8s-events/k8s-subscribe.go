package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	restclient "k8s.io/client-go/rest"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"os"
)

func main(){
	user := os.Args[1]
	server := os.Args[2]
	port := os.Args[3]

	config := &restclient.Config{
		Host:     "https://"+server+":"+port,
		Username: user,
		//Password: "supersecretpw",
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Client not created sucessfully:", err)
	}
	//Create a cache to store Pods
	var podsStore cache.Store
	var svcsStore cache.Store
	var nodesStore cache.Store
	//Watch for Pods
	podsStore = watchPods(kubeClient, podsStore)
	svcsStore = watchSvcs(kubeClient, svcsStore)
	nodesStore = watchNodes(kubeClient, nodesStore)
	//Keep alive
	log.Fatal(http.ListenAndServe(":9191", nil))
}
func podCreated(obj interface{}) {
	pod := obj.(*api.Pod)
	fmt.Println("Pod created: "+pod.ObjectMeta.Name)
}
func podDeleted(obj interface{}) {
	pod := obj.(*api.Pod)
	fmt.Println("Pod deleted: "+pod.ObjectMeta.Name)
}
func watchPods(client *kubernetes.Clientset, store cache.Store) cache.Store {
	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(client.Core().RESTClient(), "pods", api.NamespaceAll, fields.Everything())
	resyncPeriod := 30 * time.Minute
	//Setup an informer to call functions when the watchlist changes
	eStore, eController := cache.NewInformer(
		watchlist,
		&api.Pod{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    podCreated,
			DeleteFunc: podDeleted,
		},
	)
	//Run the controller as a goroutine
	go eController.Run(wait.NeverStop)
	return eStore
}
func svcCreated(obj interface{}) {
	svc := obj.(*api.Service)
	fmt.Println("Service created: "+svc.ObjectMeta.Name)
}
func svcDeleted(obj interface{}) {
	svc := obj.(*api.Service)
	fmt.Println("Service deleted: "+svc.ObjectMeta.Name)
}
func watchSvcs(client *kubernetes.Clientset, store cache.Store) cache.Store {
	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(client.Core().RESTClient(), "services", api.NamespaceAll, fields.Everything())
	resyncPeriod := 30 * time.Minute
	//Setup an informer to call functions when the watchlist changes
	eStore, eController := cache.NewInformer(
		watchlist,
		&api.Service{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    svcCreated,
			DeleteFunc: svcDeleted,
		},
	)
	//Run the controller as a goroutine
	go eController.Run(wait.NeverStop)
	return eStore
}
func nodeCreated(obj interface{}) {
	node := obj.(*api.Node)
	fmt.Println("Node created: "+node.ObjectMeta.Name)
}
func nodeDeleted(obj interface{}) {
	node := obj.(*api.Node)
	fmt.Println("Node deleted: "+node.ObjectMeta.Name)
}
func watchNodes(client *kubernetes.Clientset, store cache.Store) cache.Store {
	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(client.Core().RESTClient(), "nodes", api.NamespaceAll, fields.Everything())
	resyncPeriod := 30 * time.Minute
	//Setup an informer to call functions when the watchlist changes
	eStore, eController := cache.NewInformer(
		watchlist,
		&api.Node{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    nodeCreated,
			DeleteFunc: nodeDeleted,
		},
	)
	//Run the controller as a goroutine
	go eController.Run(wait.NeverStop)
	return eStore
}