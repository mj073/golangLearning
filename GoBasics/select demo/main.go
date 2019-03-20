package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("in fibonacci func")
	x, y := 0, 1
	fmt.Printf("in fibonacci func..x is %d;y is %d\n", x,y)
	for {
		fmt.Println("in fibonacci func's for loop ")
	select {
		case c <- x:
			fmt.Println("in fibonacci func's case c <- x")
			x, y = y, x+y
			fmt.Printf("in fibonacci func's case c <- x...x is %d;y is %d\n", x,y)
		case <-quit:
			fmt.Println("in fibonacci func's case <-quit")
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	fmt.Println("in main")
	fmt.Println("before anonymous func")
	go func() {
		fmt.Println("in anonymous func")
		for i := 0; i < 10; i++ {
			fmt.Println("i is ",i)
			fmt.Println("c is ",c)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fmt.Println("before fibonacci func call")
	fibonacci(c, quit)
}