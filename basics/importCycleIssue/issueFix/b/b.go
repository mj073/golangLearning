package b

import (
	"fmt"

	"basics/importCycleIssue/issueFix/i"
)

type B struct {
	o i.Ageter
}

func (b B) PrintB() {
	fmt.Println(b)
}

func NewB(o i.Ageter) *B {
	b := &B{o: o}
	return b
}

func RequireA(o i.Aprinter) {
	o.PrintA()
}

func (b *B) UseC() {
	fmt.Println("need to use B:",b.o.GetA().C)
}