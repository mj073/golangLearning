package main

import (
	"os/exec"
	"os"
	"log"
	"fmt"
)

func main() {
	cmd_str := "docker"
	cmd_args := []string{"-H","tcp://192.168.0.118:2376", "ps" ,"-aq", "|","xargs","docker","-H","tcp://192.168.0.118:2376", "inspect" ,"-f","'{{.Config.Hostname}} {{.Node.Name}}{{.Name}} {{.State.Status}}' "}
	cmd := exec.Command(cmd_str,cmd_args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("error while running the command:",err)
	}
	fmt.Println(cmd.Stdout)

}