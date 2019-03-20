package main

import (
"fmt"
)
type Shape interface{
	PrintArea()
	PrintCircumference()
}

type Circle struct{
	r int
}
type Square struct{
	l int
}
type Rectangle struct{
	l int
	b int
}

func (c Circle) PrintArea(){
	fmt.Println(3 * c.r * c.r)
}

func (s Square) PrintArea(){
	fmt.Println(s.l * s.l)
}

func (r Rectangle) PrintArea(){
	fmt.Println(r.l * r.b)
}

func (c Circle) PrintCircumference(){
	fmt.Println(2 * 3 * c.r)
}

func (s Square) PrintCircumference(){
	fmt.Println(4 * s.l)
}

func (r Rectangle) PrintCircumference(){
	fmt.Println(2 * r.l + 2*r.b )
}
func main() {
	fmt.Println("Hello, playground")
	var s interface{}
	s.(Circle).r = 10
	s.(Circle).PrintArea()
	//s = Square{l:10}
	//s.PrintArea()
}

