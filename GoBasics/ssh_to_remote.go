package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func main(){
	remote := "172.17.2.110"
	clientConfig := &ssh.ClientConfig{
		User: "nvmf",
		Auth: []ssh.AuthMethod{
			ssh.Password("nvmf"),
		},
	}
	if clientConfig == nil {
		fmt.Println("clientConfig is nil...")
		return
	}
	client, err := ssh.Dial("tcp", remote+":22", clientConfig)
	if err != nil {
		fmt.Println("Failed to dial: " + err.Error())
		return
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	defer session.Close()
	//go func(){
	//	w,_ := session.StdinPipe()
	//	defer w.Close()
	//	//fmt.Fprintln(w, "D0755", 0, "/tmp/records_2") // mkdir
	//	//fmt.Fprintln(w,"C0644", len(bytes), "ipdr.txt.tar.gz")
	//	//fmt.Fprint(w, string(bytes))
	//	fmt.Fprint(w, "\x00")
	//}()
	cmd := "bash -c \"hostname;echo;ls -ltrh\""
	//if err := session.Run(cmd); err != nil {
	//	panic("Failed to run: " + err.Error())
	//}
	bytes,err := session.Output(cmd)
	if err != nil{
		panic("failed to execute the remote cmd")
	}
	fmt.Println("output:",string(bytes))


}
