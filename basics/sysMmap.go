package main

import (
	"unsafe"
	"syscall"
	"fmt"
)
var (
	BasePointer = basePointer()
	BaseAddress = uintptr(BasePointer)
)
func init(){
	fmt.Printf("BasePointer:%v,BaseAddress:%x\n",BasePointer,BaseAddress)
}
func basePointer() unsafe.Pointer {
	// ok for all 32 bit devices.
	x, err := syscall.Mmap(0, 0, 1<<32, syscall.PROT_READ, syscall.MAP_PRIVATE|syscall.MAP_ANON|syscall.MAP_NORESERVE)
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(&x[0])
}

func main(){

}
