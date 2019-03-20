package i

import (
	"fmt"
)

type Aprinter interface {
	PrintA()
}
type Ageter interface {
	GetA() *A
}
type A struct {
	B *B
	C *C
}
type B struct {
	o Ageter
}
type C struct {
	o Ageter
}
func (b B) PrintB() {
	fmt.Println(b)
}

func NewB(o Ageter) *B {
	b := &B{o: o}
	return b
}


func (b *B) UseC() {
	fmt.Println("need to use C:",b.o.GetA().C)
}
func (b C) PrintB() {
	fmt.Println(b)
}

func NewC(o Ageter) *C {
	c := &C{o: o}
	return c
}

func RequireA(o Aprinter) {
	o.PrintA()
}

func (c *C) UseB() {
	fmt.Println("need to use B:",c.o.GetA().B)
}