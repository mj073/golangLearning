package main

import "container/list"

func main() {
	s := list.New()
	s.PushBack(10)
	s.PushBack(104)
	println(s.Remove(s.Back()))
	println(s.Remove(s.Back()))
}
