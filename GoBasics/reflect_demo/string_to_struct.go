package main

import (
	"os"
	"fmt"
)
type Interface struct {
	name string
	ip4 string
	mac string
	mtu int
}
func (i Interface) ping(host string){
	fmt.Println("ping host from ",i.name)
}
var Ifs [3]Interface
func init(){
	Ifs = [...]Interface{
		Interface{
			name: "eth-1-1",
			ip4: "192.168.56.101",
			mac: "00:00:00:00:00:aa",
			mtu: 1500,
		},
		Interface{
			name: "eth-2-1",
			ip4: "192.168.56.102",
			mac: "00:00:00:00:00:ab",
			mtu: 1500,
		},
		Interface{
			name: "eth-3-1",
			ip4: "192.168.56.103",
			mac: "00:00:00:00:00:ac",
			mtu: 1500,
		},
	}
}
func main() {
	v := os.Args[1]
	fmt.Println("v:",v)
	var x Interface
	ParseInput(v,&x)
}

func ParseInput(in string,x interface{}){
	switch x.(type) {
	case *Interface:
		fmt.Println("type is *Interface")
	default:
		fmt.Println("type is not *Interface")
	}
}
