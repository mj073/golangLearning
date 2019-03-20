package main

import (
	"fmt"
)

func main() {
	success := true
	mem_records := []int{1,2,3,4,5,6,7,8,9}
	for ;;{
		fmt.Println("before len:",len(mem_records))
		fmt.Println("before mem_records:",mem_records)
		for i,mem_record := range(mem_records){
			if mem_record % 4 == 0 {
				fmt.Println("condition matched")
				success = false
				break
			}
			mem_records = mem_records[i+1:]
		}
		if success == true{
			mem_records = nil
		}
		fmt.Println("after len:",len(mem_records))
		fmt.Println("after mem_records:",mem_records)
	}
}