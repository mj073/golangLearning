package main

import(
	"fmt"
	"os"
)
func main(){
	fmt.Println("PATH=",os.Getenv("PATH"))
}
