package main

import (
	"fmt"
	"time"
	"os/exec"
)

func main(){
	getEventVariable("main:")
}
func getEventVariable(lp string){

	//args := "source /home/mahesh/environment && echo -n $Event"
	//args := "set Event=true && echo -n $Event"
	out,err := exec.Command("bash","/home/mahesh/getEvent.sh").Output()
	if err != nil{
		fmt.Println(lp,"failed to get event variable..ERROR:",err)
	}else if string(out)== "true" {
		fmt.Println(lp, "sleeping for 1 minute...")
		time.Sleep(time.Minute * 1)
	}else{
		fmt.Println("out value is not true...out=",string(out))
	}
}
