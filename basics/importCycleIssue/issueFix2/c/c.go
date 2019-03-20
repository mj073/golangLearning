package c

import (
	"fmt"

	"basics/importCycleIssue/issueFix2/types"
)

type C types.C


func NewC(o types.Ageter) *C {
	c := &C{O: o}
	return c
}

func (c *C) UseB() {
	fmt.Println("need to use B:",c.O.GetA().B)
}