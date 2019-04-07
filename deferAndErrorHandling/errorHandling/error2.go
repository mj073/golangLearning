package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f,err := os.OpenFile("/home/mahesh/golang/src/github.com/golangLearning/deferAndErrorHandling/defer/defer2.go",os.O_RDONLY,0644)
	if err != nil {
		log.Fatalln("failed to open file...ERROR:",err)
	}
	defer f.Close()
	data := make([]byte,450)
	_, err = f.Read(data)
	if err != nil{
		log.Fatalln("failed to read file")
	}
	fmt.Println(string(data))
}