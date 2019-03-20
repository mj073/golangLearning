package main

import "fmt"

func main(){
	arr := []int{9,4,8,8,2,5,87,43,25,3,6,8,7,7}
	out := make([]int,len(arr))
	mergeSort(len(arr),arr,out)
	fmt.Printf("sorted array arr: %v",out)
}
func mergeSort(n int,a []int,b []int){
	if n<2{
		copy(b,a)
	}else {
		a1 := make([]int, n / 2)
		a2 := make([]int, n - n / 2)
		mergeSort(n / 2, a[0:n / 2], a1)
		mergeSort(n - n / 2, a[n / 2 :], a2)
		merge(n/2,a1,n-n/2,a2,b)
	}
}

func merge(n1 int,a1 []int,n2 int,a2 []int,b []int){
	i1, i2, ib := 0,0,0
	for (i1 < n1) || (i2 < n2){
		if i2 >= n2 || ((i1 < n1) && (a1[i1] < a2[i2])){
			b[ib] = a1[i1]
			i1++
		}else {
			b[ib] = a2[i2]
			i2++
		}
		ib++
	}
}