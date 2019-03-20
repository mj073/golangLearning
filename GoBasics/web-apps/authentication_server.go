package main

import (
	"net/http"
	"fmt"
	"log"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/validate",validate)
	err := http.ListenAndServe(":8888",mux)
	if err != nil{
		fmt.Println("error while starting go server...ERROR:",err)
	}
}

func validate(w http.ResponseWriter,r *http.Request){
	buff := []byte{}

	n,err := r.Body.Read(buff)
	if err != nil{
		log.Println("error while reading request body..ERROR:",err)
		fmt.Fprintln(w,"http_status_code:",http.StatusInternalServerError,"ERROR:"+err.Error())
		return
	}
	log.Println("bytes read from request body:",n)
	log.Println("request body details: ",string(buff))
}