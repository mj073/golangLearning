package main

import (
	"fmt"
	"math"
	"runtime/trace"
	"os"
	"errors"
)

const Pie = math.Pi

func main() {
	f, err := os.Open("basics/trace.out")
	if err != nil {
		panic(errors.New("failed to open file...ERROR:"+err.Error()))
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("original arr before any modifications: \t\t\t\t\t\t", arr)
	modifyArrayPassByValue(arr)
	fmt.Println("original arr after modification in pass by value func: \t\t\t", arr)
	modifyArrayPassByReference(&arr)
	fmt.Println("original arr after modification in pass by reference func: \t\t\t", arr)
	fmt.Println(Pie)
}

func modifyArrayPassByValue(arr [10]int) {
	arr[5] = 100
	fmt.Println("modifyArrayPassByValue: modified arr: \t\t\t\t\t", arr)
}

func modifyArrayPassByReference(arr *[10]int) {
	arr[5] = 100
	fmt.Println("modifyArrayPassByReference: modified arr: \t\t\t\t", *arr)
}
