package main

import (
	"os/exec"
	"fmt"
	"io"
	"bytes"
	"strings"

)

func main(){
	arg := "docker -H tcp://192.168.0.118:2376 ps -aq | xargs -I ID -P 0 bash /home/alef/docker-monitor.sh ID 192.168.0.118 | sed 's/,$//g'"
	out,err := exec.Command("bash","-c",arg).Output()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("command success")
	fmt.Println(string(out))

	//xargs_args := []string{"-I","ID","-P","0","bash","/home/alef/docker-monitor.sh","ID","192.168.0.118"}
	//echo_first := exec.Command("echo","-n","[")
	//lscpu_second := exec.Command("docker","-H","tcp://192.168.0.118:2376","ps","-aq")
	//lscpu_third := exec.Command("xargs",xargs_args...)
	////lscpu_fourth := exec.Command("sed","'s/,$//g'")
	//echo_last := exec.Command("echo","-n","]")
	//
	////lscpu_second := exec.Command("seq","12")
	////lscpu_third := exec.Command("xargs","-I","ID","echo","ID")
	////lscpu_fourth := exec.Command("sed","'s/.$//g'")
	//
	//final := Pipeline(echo_first) + strings.TrimSuffix(Pipeline(lscpu_second,lscpu_third),",") + Pipeline(echo_last)
	//
	////final := Pipeline(lscpu_second,lscpu_third)
	//fmt.Println(final)
}

func Pipeline(cmds...*exec.Cmd) string{
	fmt.Println("in Pipeline function")
	var reader = make([]*io.PipeReader,len(cmds)-1)
	var writer = make([]*io.PipeWriter,len(cmds)-1)
	for i := 0;i<len(cmds)-1;i++ {
		reader[i],writer[i] = io.Pipe()
	}
	var out bytes.Buffer

	for i := 0;i<len(cmds)-1;i++{
		cmds[i].Stdout = writer[i]
		cmds[i+1].Stdin = reader[i]
	}
	cmds[len(cmds)-1].Stdout = &out
	for i := 0;i<len(cmds);i++{
		fmt.Println("starting cmd:",cmds[i].Path," arguments:",cmds[i].Args)
		cmds[i].Start()
	}

	for i := 0; i < len(cmds)-1; i++{
		fmt.Println("i=",i)
		cmds[i].Wait()
		writer[i].Close()
	}
	cmds[len(cmds)-1].Wait()

	final_output := strings.TrimSuffix(out.String(),"\n")
	//fmt.Println("output in bytes:",out)
	//fmt.Println("final output:\n",final_output)

	return final_output
}