package main

import (
	"log"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"fmt"
	"time"
)
func main() {
	remote := "192.168.0.118"
	file := "/home/alef/ipdr.txt.tar.gz"
	bytes,_ := ioutil.ReadFile(file)
	priv_key, err := ioutil.ReadFile("/root/temp_id_rsa")
	if err != nil{
		log.Fatalln("error while reading private key..ERROR:",err)
	}
	signer,err := ssh.ParsePrivateKey(priv_key)
	if err != nil{
		log.Fatalln("error while parsing private key..ERROR:",err)
	}
	clientConfig := &ssh.ClientConfig{
		User: "mjadhav",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout: time.Minute * 1,
	}
	//clientConfig := &ssh.ClientConfig{
	//	User: "alef",
	//	Auth: []ssh.AuthMethod{
	//		ssh.Password("alef@123"),
	//	},
	//}
	if clientConfig == nil {
		log.Fatalln("clientConfig is nil...")
	}
	client, err := ssh.Dial("tcp", remote+":22", clientConfig)
	if err != nil {
		log.Fatalln("Failed to dial: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	go func(){
		w,_ := session.StdinPipe()
		defer w.Close()
		//fmt.Fprintln(w, "D0755", 0, "/tmp/records_2") // mkdir
		fmt.Fprintln(w,"C0644", len(bytes), "ipdr.txt.tar.gz")
		fmt.Fprint(w, string(bytes))
		fmt.Fprint(w, "\x00")
	}()

	//go func() {
	//	w, _ := session.StdinPipe()
	//	defer w.Close()
	//	content := "123456789\n"
	//	fmt.Fprintln(w, "D0755", 0, "testdir") // mkdir
	//	fmt.Fprintln(w, "C0644", len(content), "testfile1")
	//	fmt.Fprint(w, content)
	//	fmt.Fprint(w, "\x00") // transfer end with \x00
	//	fmt.Fprintln(w, "C0644", len(content), "testfile2")
	//	fmt.Fprint(w, content)
	//	fmt.Fprint(w, "\x00")
	//}()
	if err := session.Run("/usr/bin/sftp -tr /tmp/"); err != nil {
		panic("Failed to run: " + err.Error())
	}
}
