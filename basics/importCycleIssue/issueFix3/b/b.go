package b

import (
	"fmt"

	"basics/importCycleIssue/issueFix3/i"
)

type B struct {
	a i.A
}

func NewB(a i.A) *B {
	b := &B{a: a}
	return b
}

func (b *B) UseC() {
	fmt.Println("need to use C:",b.a.GetC())
}