package c

import (
	"fmt"

	"basics/importCycleIssue/issueFix/i"
)

type C struct {
	o i.Ageter
}

func (b C) PrintB() {
	fmt.Println(b)
}

func NewC(o i.Ageter) *C {
	c := &C{o: o}
	return c
}

func RequireA(o i.Aprinter) {
	o.PrintA()
}

func (c *C) UseB() {
	fmt.Println("need to use B:",c.o.GetA().B)
}