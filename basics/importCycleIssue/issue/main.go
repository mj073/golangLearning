package main

import (
	"basics/importCycleIssue/issue/a"
	"basics/importCycleIssue/issue/b"
)

func main(){
	o := a.NewA()
	b.RequireA(o)
}
