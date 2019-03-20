package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000.")
	flag.Parse()

	fmt.Printf("port = %d", port)

	fmt.Println("bad gateway: ",http.StatusInternalServerError)
}