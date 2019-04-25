package main

import (
	"fmt"
	"os"
)

func main() {
	f,err := os.OpenFile("/home/mahesh/golang/src/github.com/golangLearning/golangTraining/deferAndErrorHandling/defer/defer2.go",os.O_RDONLY,0644)
	if err != nil {
		fmt.Println("failed to open file...ERROR:",err)
		os.Exit(1)
	}
	defer f.Close()
	data := make([]byte,450)
	_, err = f.Read(data)
	if err != nil{
		fmt.Println("failed to read file")
		os.Exit(1)
	}
	fmt.Println(string(data))
}