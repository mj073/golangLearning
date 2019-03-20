package main

import (
	"time"
	"fmt"
	"strconv"
)
const (
	custom_time_format = "2006-01-02 15:04:05"
)
func main(){
	//t := time.Time{}

	timeInUnix := "1470734979"
	timeNow := time.Unix(time.Now().Unix(),0)
	i, _ := strconv.ParseInt(timeInUnix,10,64)
	t := time.Unix(i,0)
	if t.After(timeNow){
		fmt.Println("true: ",t.String(),"is after ",timeNow.String())
	}else {
		fmt.Println("false: ",timeNow.String(),"is after ",t.String())
	}
	fmt.Println(t.String())
	fmt.Print(t.Format(custom_time_format))
}