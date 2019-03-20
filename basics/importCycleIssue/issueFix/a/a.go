package a

import (
	"fmt"

	"basics/importCycleIssue/issueFix/b"
	"basics/importCycleIssue/issueFix/c"
)

type A struct {
	B *b.B
	C *c.C
}
var aa *A
func (a A) PrintA() {
	fmt.Println(a)
}

func NewA() *A {
	a := &A{}
	a.B = b.NewB(a)
	a.C = c.NewC(a)
	aa = a
	return a
}
func GetA() *A{
	return aa
}

func RequireB() {
	o := b.NewB(aa)
	o.PrintB()
}