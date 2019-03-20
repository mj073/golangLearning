//An Armstrong Number is a Number which is equal to it’s sum of digit’s cube.
//For example - 153 is an Armstrong number: here 153 = (1*1*1) + (5*5*5) + (3*3*3).
package main

import "fmt"

func main(){
	num := 153
	temp := num
	rem := -1
	sum := 0
	for rem != 0{
		rem = temp % 10
		//fmt.Println("rem:",rem)
		temp /= 10
		sum = sum + rem * rem * rem
		//fmt.Println("sum:",sum)
	}
	if sum == num{
		fmt.Println(num,"is an Armstrong Number..")
	}else {
		fmt.Println(num,"is not an Armstrong Number")
	}
}
