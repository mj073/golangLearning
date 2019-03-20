package main

import "fmt"

func main(){
	arr := []int{9,4,8,8,2,5,87,43,25,3,6,8,7,7}

	for i:=0; i<len(arr); i++{
		min := arr[i]
		min_index := i
		for j:=i+1;j<len(arr);j++{
			if arr[j] < min{
				min = arr[j]
				min_index = j
			}
		}
		if min_index != i{
			//temp := arr[i]
			arr[i], arr[min_index] = arr[min_index], arr[i]
			//arr[min_index] = temp
		}
	}
	fmt.Printf("sorted array arr: %v",arr)
}
