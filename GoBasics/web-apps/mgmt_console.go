package main

import (
	"net/http"
	"fmt"
	"log"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/register",register)
	err := http.ListenAndServe(":9999",mux)
	if err != nil{
		fmt.Println("error while starting go server...ERROR:",err)
	}
}

func register(w http.ResponseWriter,r *http.Request){
	buff := []byte{}

	n,err := r.Body.Read(buff)
	if err != nil{
		log.Println("error while reading request body..ERROR:",err)
		fmt.Fprintln(w,"http_status_code:",http.StatusInternalServerError,"ERROR:"+err.Error())
		return
	}
	log.Println("bytes read from request body:",n)
	log.Println("request body details: ",string(buff))
	fmt.Fprintln(w,"request body:",string(buff))
}
