package main

import (
	"bytes"
	"fmt"
)

func main() {
	msg := "hello"
	//buf := bytes.NewBufferString(msg)
	bufr := bytes.NewReader([]byte(msg))
	//bufr := bufio.NewReader(buf)
	bytes, _ := bufr.Read(nil)
	fmt.Println(bytes)
}
