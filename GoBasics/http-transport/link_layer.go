package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_UDP)
	if err != nil{
		fmt.Println("error while opening socket..ERROR:",err)
		return
	}
	fmt.Println("socket opened..",fd)
	file := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))
	defer file.Close()

	for {
		buf := make([]byte,2048)
		if file == nil{
			fmt.Println("f is nil")
			//break
		}
		numRead, err := file.Read(buf)
		if err != nil {
			fmt.Println("failed to read the buffer..ERROR:",err)
			//break
		}
		fmt.Printf("%d \n", buf[:numRead])
	}
}