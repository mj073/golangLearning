package a

import (

	"basics/importCycleIssue/issueFix2/b"
	"basics/importCycleIssue/issueFix2/c"
	"basics/importCycleIssue/issueFix2/types"
)

type A types.A
var aa *A
func NewA() *types.A {
	a := &A{}
	a.B = b.NewB(a)
	a.C = c.NewC(a)
	aa = a
	return a
}
func GetA() *types.A{
	return aa
}

func RequireB() {
	o := b.NewB(aa)
	o.PrintB()
}