package main

import (
	"fmt"
	"time"
)
type Person struct {
	Name	string
	Age chan int
}

var sem = make(chan int,1)
func main(){
	p := Person{
		Name: "Mahesh",
	}
	p.Age = make(chan int,10)
	go p.Run()
	go func() {
		c := 0
		for c < 10{
			c += 1
			p.Age <- c
			time.Sleep(time.Second)
		}
		close(p.Age)
	}()

}

//func Random(i int){
//	log.Println("in random..i:",i)
//	sem <- i
//	log.Println("in Random..sem:",sem)
//	<- sem
//}

func (p Person) Run(){
	fmt.Println("in run")
	for{
		switch <-p.Age{
		case 5:
			fmt.Println(p.Name," you are 5 seconds old..!!")
		case 10:
			fmt.Println(p.Name," you are 10 seconds old..!!")
		}
	}
}