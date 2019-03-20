package main

import (
	"basics/importCycleIssue/issueFix3/a"
)

func main() {
	o := a.NewA()
	o.B.UseC()
	o.C.UseB()
}