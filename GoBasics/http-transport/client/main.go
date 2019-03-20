package main

import "net/http"

func main(){
	req,err := http.NewRequest("GET","/containers/*/stats")
}