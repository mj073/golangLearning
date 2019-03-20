package a

import (

	"basics/importCycleIssue/issueFix3/b"
	"basics/importCycleIssue/issueFix3/c"
	"basics/importCycleIssue/issueFix3/i"
)

type A struct {
	B *b.B
	C *c.C
}

func NewA() *A {
	a := &A{}
	a.B = b.NewB(a)
	a.C = c.NewC(a)
	return a
}

// These methods implement i.A and return the i.B and i.C interface types
func (a A) GetB() i.B {
	return a.B
}

func (a A) GetC() i.C {
	return a.C
}