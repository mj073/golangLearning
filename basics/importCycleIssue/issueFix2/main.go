package main

import (
	"basics/importCycleIssue/issueFix2/a"
)

func main() {
	o := a.NewA()
	o.B.UseC()
	o.C.UseB()
}