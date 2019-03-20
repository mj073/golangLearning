package main

import (
	"fmt"
	"net/http"
	"os"
	//"encoding/json"
	//"GoBasics/json-parsing-demo/handlers"
	//"encoding/json"
)
type Name struct {
	Id string
}
func main(){
	fmt.Println("in client's main")
	resp,err := http.Get("http://192.168.0.108:9000/get/Mahesh")
	if err != nil {
		fmt.Println("error occurred:",err)
		os.Exit(0)
	}
	//name := Name{}
	contents := make([]byte,1000)
	fmt.Println(resp.Body.Read(contents))
	fmt.Println(string(contents))
}
