package main

import (
	"fmt"
	"time"
	"encoding/json"
	"reflect"
	"net/http"
	"strconv"
)

type Meta struct {
	Name string
	Age int
	Birthdate time.Time
}
func main() {

	meta := []Meta{}
	fmt.Println("len(meta):",len(meta))
	fmt.Println(meta)
	//append(meta, Meta{"Mahesh",20,time.Unix(1467013837,0)})
	for i,_ := range meta{
		fmt.Println("meta:   ",meta[i])
	}
	fmt.Println(time.Unix(1467013837,0))
	person := Meta{"Mahesh",20,time.Unix(1467013837,0)}
	s := reflect.ValueOf(&person).Elem()
	typeOfT := s.Type()
	fmt.Println("for loop begin")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	fmt.Println("for loop end")
	jsonTime,_ := json.Marshal(person)
	unmarshalled := Meta{}
	json.Unmarshal(jsonTime,&unmarshalled)
	fmt.Println(unmarshalled)

	files := make([]string,10)
	for i:=0;i<5;i++{
		files[i] = "hello"
	}
	if files[5] == "" {
		fmt.Println(reflect.TypeOf(files[5]))
	}
	//randomsource := "c:/dev/randomfile.txt"
	//_, err := os.Open(randomsource)
	//if err != nil {
	//	log.Fatal("cannot open random source:", err)
	//}
	fmt.Println("time.Now():",time.Now())
	threshold := time.Now().Add(- (time.Second * 20))
	fmt.Println("threshold:",threshold)
	fmt.Println(strconv.Itoa(http.StatusOK))
}


