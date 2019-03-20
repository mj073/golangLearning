package main

import (
	"os/exec"
	"fmt"
)

func main(){
	cmdargs := "ifconfig | grep \"10.0.2.15\""
	output,_ := exec.Command("bash","-c",cmdargs).Output()

	fmt.Println(string(output))

}