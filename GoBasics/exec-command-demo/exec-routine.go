package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main(){
	//cmdargs := "sudo ifconfig|grep 172.17.2.25"
	//cmdargs := "sudo ls -ltrh|grep data.origin"
	cmdargs1 := "sudo ./bcm.py \"listreg CLMAC_PAUSE_CTRL\""
	cmdargs2 := "sudo ./bcm.py \"listreg XLMAC_PAUSE_CTRL\""
	cmdargs3 := "sudo ./bcm.py \"listreg MMU_THDM_MCQE_DEVICE_THR_CONFIG\""

	go execute(cmdargs1)
	go execute(cmdargs2)
	go execute(cmdargs3)

	time.Sleep(time.Minute * 1)
}

func execute(cmd string){
	fmt.Println("cmd:",cmd)

	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("output:",string(output))
	fmt.Println("--------------------------")
}