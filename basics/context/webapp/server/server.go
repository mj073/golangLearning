package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/search",handleSeatch)
	if err := http.ListenAndServe(":8080",nil); err != nil{
		panic(err)
	}
}

func handleSeatch(w http.ResponseWriter, req *http.Request){

	timeout := req.FormValue("timeout")
	fmt.Print("timeout:",timeout)
	w.Write([]byte("found"))
}