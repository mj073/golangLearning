package main

import (
	"fmt"
)
type PodByName map[string]string
type ContainersById map[string]string
type Resource struct{
	podByName PodByName
	containersById ContainersById
	containersByPod map[string]ContainersById
}
type ResourceLookup struct{

}
func main() {
	fmt.Println("Hello, playground")
	resource := &Resource{
		podByName: PodByName{
			"pod1": "pod1",
			"pod2": "pod2",
		},
		containersById: ContainersById{
			"1": "cont1",
			"2": "cont2",
		},
		containersByPod: make(map[string]ContainersById),
	}
	resource.containersByPod["pod1"] = make(ContainersById)
	resource.containersByPod["pod1"]["1"] = resource.containersById["1"]
	fmt.Println(resource)
	delete(resource.containersById,"1")
	fmt.Println(resource)
}


