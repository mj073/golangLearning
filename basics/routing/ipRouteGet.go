package main

import (
	"time"
	"fmt"
	"os"
	"basics/routing/lib"
)

func main(){
	src := os.Args[1]
	dst := os.Args[2]
	go lib.UpdateRoutingTableInfo()
	<- time.After(time.Second * 1)
	fmt.Println(lib.GetInterface(src,dst))
	<- time.After(time.Minute * 3)
}
