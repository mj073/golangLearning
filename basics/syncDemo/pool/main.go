package main

import (
	"sync"
	"fmt"
)
func main(){

	p := &sync.Pool{
		New: func() interface{}{
			return make([]byte, 4, 5)
		},
	}

	b := []byte{1,4,91}
	p.Put(b)
	y := p.Get().([]byte)
	fmt.Println(y,len(y))

}
