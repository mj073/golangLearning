package main

import (
	oauth "golang.org/x/oauth2"
)

func main(){
	oauth.RegisterBrokenAuthHeaderProvider("http://192.168.0.108:8000/authenticate")

}

