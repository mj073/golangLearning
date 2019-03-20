package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello, playground")
	x := new([]string)

	fmt.Println("x:",&x)
	fmt.Println("unsafe.Prointer(x):",unsafe.Pointer(x))
	uint64FromSystem := (uintptr)(unsafe.Pointer(x))
	fmt.Println("uint64FromSystem :",uint64FromSystem)
}
