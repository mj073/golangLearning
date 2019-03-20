package main

import (
	"net/http"
	"log"
	"GoBasics/json-parsing-demo/handlers"
)


func main() {

	r := http.NewServeMux()
	//r.HandleFunc("/get/",handlers.GetHandler)
	//r.HandleFunc("/save/",handlers.SaveHandler)
	r.HandleFunc("/get",handlers.GetHandler)
	log.Fatal(http.ListenAndServe("192.168.0.11:9000",r))

}