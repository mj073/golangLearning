//Perfect number, a positive integer that is equal to the sum of its proper divisors.
// The smallest perfect number is 6, which is the sum of 1, 2, and 3.
// Other perfect numbers are 28, 496, and 8,128

package main

import (
	"fmt"
)

func main(){
	var num int
	fmt.Print("enter any number:")
	fmt.Scanf("%d",&num)
	sum := 0
	for i := num/2;(i>=1 )&& (num %i == 0);i--{
		sum = sum + i
	}
	if sum == num{
		fmt.Printf("%d is a perfect number",num)
	}else{
		fmt.Printf("%d is not a perfect number",num)
	}
}
