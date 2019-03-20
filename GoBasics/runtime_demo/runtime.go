package main

import (
	"runtime"
	"log"
)
var b = make([]byte,3000)
var random chan int
func main(){
	//var buf bytes.Buffer

	log.Println("GOOS:",runtime.GOOS)
	log.Println("GOARCH:",runtime.GOARCH)
	log.Println("Compiler:",runtime.Compiler)
	log.Println("CPUProfile:",runtime.CPUProfile())
	log.Println("ReadTrace:",string(runtime.ReadTrace()))
	log.Println("NumCgoCall:",runtime.NumCgoCall())
	log.Println("MemProfileRate:",runtime.MemProfileRate)
	log.Println("Stack:",runtime.Stack(b,false))
	log.Println("b:",string(b))
	go greetings()
	<- random
}
func greetings(){
	//b := make([]byte,3000)
	log.Println("greetings")
	log.Println("Stack:",runtime.Stack(b,false))
	log.Println("b:",string(b))
	hello()
	log.Println("greetings_hello")
	log.Println("Stack:",runtime.Stack(b,false))
	log.Println("b:",string(b))
	bye()
}
func bye(){
	//b := make([]byte,3000)
	log.Println("bye bye world")
	log.Println("Stack:",runtime.Stack(b,false))
	log.Println("b:",string(b))
}
func hello(){
	//b := make([]byte,3000)
	log.Println("hello world")
	log.Println("Stack:",runtime.Stack(b,false))
	log.Println("b:",string(b))
}
