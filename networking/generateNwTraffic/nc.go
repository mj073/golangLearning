package main

import (
	"os/exec"
	"fmt"
	"time"
)

func main(){
	for i :=0; i < 30000;i++{
		cmd := exec.Command("nc","-vz","192.168.5.32","179")
		//fmt.Println("cmd:",cmd)
		go func (){
			err := cmd.Run()
			if err != nil {
				fmt.Println("error:", err)
			}
		}()
	}
	time.Sleep(time.Second * 50)
}
