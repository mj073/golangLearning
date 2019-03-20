package main

import (
	api "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	"time"
	"log"
	"net/http"
	"os"
	"fmt"
)

var Nodes []*api.Node
var Services []*api.Service
var Pods []*api.Pod

var NodeIndex int
var ServiceIndex int
var PodIndex int

var NodeIndexByName = make(map[string]int)
var ServiceIndexByName = make(map[string]int)
var PodIndexByName = make(map[string]int)

var NodeByName = make(map[string]*api.Node)
var ServiceByName = make(map[string]*api.Service)
var PodByName = make(map[string]*api.Pod)

var NodeByIp = make(map[string]*api.Node)
var ServiceByIp = make(map[string]*api.Service)
var PodByIp = make(map[string]*api.Pod)

type Resource struct {
	Name string
	Store cache.Store
	Controller cache.Controller
}
func resourceCreated(obj interface{}) {
	switch x := obj.(type){
	case *api.Node:
		Nodes = append(Nodes,x)
		NodeByName[x.Name] = x
		NodeIndex++
		NodeIndexByName[x.Name] = NodeIndex-1
	case *api.Service:
		Services = append(Services,x)
		ServiceByName[x.Name] = x
		ServiceByIp[x.Spec.ClusterIP] = x
		ServiceIndex++
		ServiceIndexByName[x.Name] = ServiceIndex - 1
		fmt.Println("Cluster Name of svc",x.Name,x.ClusterName)
	case *api.Pod:
		Pods = append(Pods,x)
		PodByName[x.Name] = x
		PodByIp[x.Status.PodIP] = x
		PodIndex++
		PodIndexByName[x.Name] = PodIndex - 1
		fmt.Println("Cluster Name of pod",x.Name,x.ClusterName)
	default:

	}
}
func resourceDeleted(obj interface{}) {
	switch x := obj.(type){
	case *api.Node:
		i := NodeIndexByName[x.Name]
		Nodes = append(Nodes[:i],Nodes[i+1:]...)
		delete(NodeByName,x.Name)
		delete(NodeIndexByName,x.Name)
	case *api.Service:
		i := ServiceIndexByName[x.Name]
		Services = append(Services[:i],Services[i+1:]...)
		delete(ServiceByName,x.Name)
		delete(ServiceByIp,x.Spec.ClusterIP)
		delete(ServiceIndexByName,x.Name)
	case *api.Pod:
		i := PodIndexByName[x.Name]
		Pods = append(Pods[:i],Pods[i+1:]...)
		delete(PodByName,x.Name)
		delete(PodByIp,x.Status.PodIP)
		delete(PodIndexByName,x.Name)
	default:

	}
}
func (resource *Resource) watchResources(client *kubernetes.Clientset){
	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(client.Core().RESTClient(), resource.Name, api.NamespaceAll, fields.Everything())
	resyncPeriod := 30 * time.Minute
	//Setup an informer to call functions when the watchlist changes
	resource.Store, resource.Controller = cache.NewInformer(
		watchlist,
		&api.Node{},
		resyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    resourceCreated,
			DeleteFunc: resourceDeleted,
		},
	)
	//Run the controller as a goroutine
	go resource.Controller.Run(wait.NeverStop)
}
type KubeWatcherConfig struct {
	ApiServer string
	ApiServerPort string
	User string
}
func InitKubeWatcher(c *KubeWatcherConfig){
	config := &restclient.Config{
		Host:     "https://"+c.ApiServer+":"+c.ApiServerPort,
		//Username: c.User,
		//Password: "supersecretpw",
	}
	kubeClient, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatalln("Client not created sucessfully:", err)
	}
	for _,r := range []*Resource {
		&Resource{Name: "nodes"},
		&Resource{Name: "services"},
		&Resource{Name: "pods"},
	}{
		r.watchResources(kubeClient)
	}
}
func main(){
	config := &KubeWatcherConfig{}
	//config.User = os.Args[1]
	config.ApiServer = os.Args[1]
	config.ApiServerPort = os.Args[2]
	InitKubeWatcher(config)
	log.Fatal(http.ListenAndServe(":9191", nil))
}

