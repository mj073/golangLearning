package b

import (
	"fmt"

	"basics/importCycleIssue/issueFix2/types"
)


type B types.B
func (b *B) PrintB() {
	fmt.Println(b)
}

func NewB(o types.Ageter) *B {
	b := &B{O: o}
	return b
}

func (b *B) UseC() {
	fmt.Println("need to use B:",b.O.GetA().C)
}