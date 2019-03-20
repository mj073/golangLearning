package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)

func main() {
	signalChannel := make(chan os.Signal, 2)
	//signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM,syscall.SIGPIPE)
	signal.Ignore(syscall.SIGTERM,syscall.SIGKILL)
	go func() {
		sig := <-signalChannel
		switch sig {
			case syscall.SIGPIPE:
				fmt.Println("SIGPIPE signal captured...")
			case os.Interrupt:
				fmt.Println("interrupt signal captured")
				os.Exit(1)
			case syscall.SIGTERM:
			//handle SIGTERM
				fmt.Println("terminate signal captured..")
			case syscall.SIGKILL:
				fmt.Println("kill signal captured..")
			case syscall.SIGHUP:
				fmt.Println("SIGHUP signal captured...")


		}
	}()
	time.Sleep(time.Second * 60)
}