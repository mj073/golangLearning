package main

import (
	"net/http"
	"log"
	"GoBasics/simple-rest-server/handlers"
	"orc_bitbucket_repo/Godeps/_workspace/src/github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	//r.HandleFunc("/get/",handlers.GetHandler)
	//r.HandleFunc("/save/",handlers.SaveHandler)
	r.HandleFunc("/get/{name}",handlers.GetHandler)
	log.Fatal(http.ListenAndServe("192.168.0.108:9000",r))
}