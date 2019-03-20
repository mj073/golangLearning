package main

import (
	"fmt"
)

//go:generate gentemplate -id=temp -d Operation=+ -d Type=int,float basics/goGenerateDemo/arithmeticFuncs.tmpl
//go:generate gentemplate -id=temp1 -d Operation=+ -d Type=int basics/goGenerateDemo/arithmeticFuncs.tmpl

func main(){
	fmt.Println(Add(1,2))
}
