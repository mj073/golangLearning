package b

import (
	"fmt"

	"basics/importCycleIssue/issue/a"
)

type B struct {
	a *a.A
}

func (b B) PrintB() {
	fmt.Println(b)
}

func NewB(a *a.A) *B {
	b := &B{a: a}
	return b
}

func RequireA(o *a.A) {
	o.PrintA()
}

func (b *B) UseA() {
	b.a.PrintA()
}