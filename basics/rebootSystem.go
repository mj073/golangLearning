package main

import (
	"syscall"
	"fmt"
)
func main(){
	fmt.Println("rebooting system...")
	err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
	if err != nil {
		panic(err)
	}
}
