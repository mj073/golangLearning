package a

import (
	"fmt"

	"basics/importCycleIssue/issue/b"
)

type A struct {
	b *b.B
}
var a = NewA()
func (a A) PrintA() {
	fmt.Println(a)
}

func NewA() *A {
	a := &A{}
	a.b = b.NewB(a)
	return a
}
func GetA() *A{
	return a
}
func RequireB() {
	o := b.NewB(a)
	o.PrintB()
}