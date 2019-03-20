package main

import (
	"basics/importCycleIssue/issueFix/a"
)

func main() {
	o := a.NewA()
	o.B.UseC()
	o.C.UseB()
}