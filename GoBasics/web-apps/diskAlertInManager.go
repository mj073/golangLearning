package main

import (
	"net/http"
	"fmt"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/diskalertManager",diskAlertManagerHandler)
	err := http.ListenAndServe(":8181",mux)
	if err != nil{
		fmt.Println("error while starting go server...ERROR:",err)
	}
}
func diskAlertManagerHandler(w http.ResponseWriter, r *http.Request){
	resp,err := http.Post("http://192.168.0.25:8282/diskalertPortal",r.Header.Get("Content-Type"),r.Body)
	if err != nil{
		fmt.Println("error while calling portal's disk alert API...ERROR:",err)
	}else if resp != nil && resp.StatusCode == http.StatusOK {
		fmt.Println("successfully sent disk alert info to portal...")
	}else{
		fmt.Println("failed to send disk alert info to portal...")
	}
}