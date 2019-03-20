package c

import (
	"fmt"

	"basics/importCycleIssue/issueFix3/i"
)

type C struct {
	a i.A
}

func NewC(a i.A) *C {
	c := &C{a: a}
	return c
}

func (c *C) UseB() {
	fmt.Println("need to use B:",c.a.GetB())
}