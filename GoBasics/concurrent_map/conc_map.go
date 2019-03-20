package main

import (
	"strconv"
	"time"
	"sync"
	"fmt"
)

type EmployeeMap map[string]struct{
	Name string
	Age int
}
var lock sync.RWMutex
func main(){
	p := make(EmployeeMap)
	c := p["Calsoft"]
	c.Name = "Mahesh"
	c.Age = 25
	for i := 0; i< 100000; i++ {
		go func(age int){
			lock.Lock()
			c := p["Calsoft"]
			c.Name = "Mahesh"+ strconv.Itoa(i)
			c.Age = age
			p["Calsoft"] = c
			lock.Unlock()
		}(i)
	}
	time.Sleep(time.Second * 15)
	fmt.Println("len of map:",len(p))
}
