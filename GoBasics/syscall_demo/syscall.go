package main

import (
	"syscall"
	"log"
	"os"
)

func main(){
	printComputerName()
	log.Println("Environment:",syscall.Environ())
	log.Println("PID:",syscall.Getpid())
	os.StartProcess()
}
func printComputerName(){
	computerName,err := syscall.ComputerName()
	if err != nil{
		log.Fatalln("error while getting computername...ERROR:",err)
	}
	log.Println("ComputerName:",computerName)
}
