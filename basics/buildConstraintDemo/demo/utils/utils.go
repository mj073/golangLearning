package utils

import (
	"runtime"
	"fmt"
)

func DoSomething(){
	fmt.Println("Doing something")
	if runtime.GOOS == "windows" {
		Done()
	}
}