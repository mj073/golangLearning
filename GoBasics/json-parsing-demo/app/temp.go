package main

import (
	"encoding/json"
	"fmt"
)

type Value struct {
	key string	`json:"key"`
	value string	`json:"value"`
}
func main(){
	//k := &Value{}
	var v map[string]string
	str := "{\"key\":\"hello:34635\",\"value\":\"world\"}"
	err := json.Unmarshal([]byte(str),&v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
