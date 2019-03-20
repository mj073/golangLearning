package main

import "fmt"

func main() {
	relmMap := make(map[string][]string)

	relmMap["1"] = append(relmMap["1"],"one")
	relmMap["1"] = append(relmMap["1"],"one")

	fmt.Println(relmMap["1"])
}
