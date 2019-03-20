package main

import "fmt"

func main(){
	arr := []int{9,4,8,8,2,5,87,43,25,3,6,8,7,7}

	for i := 1;i<len(arr);i++{
		temp := arr[i]
		j := i - 1
		for (j >= 0) && (temp < arr[j]){
			arr[j+1] = arr[j]
			arr[j] = temp
			j = j - 1
		}
	}
	fmt.Printf("sorted array arr: %v",arr)
}