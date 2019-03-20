package main

import (
	"fmt"
	"os"
	"strings"
)
type Person struct{
	Name	string
	Age int
}
type Employee struct{
	P Person
	EmpId	int
	CompanyName	string
}

func (p Person) Walk(){
	fmt.Println("person walks")
}
func (p Person) Talk(){
	fmt.Println("Person talks")
}
func (e Employee) TimeEntry(){
	fmt.Println("Employee makes time entry")
}


func main(){
	emp := Employee{Person{"Mahesh",25},123,"Alef"}
	emp.P.Talk()
	defer func(){
		fmt.Println("Hello")
		str := recover()
		if st := fmt.Sprintf("%v",str);strings.Contains(st,"runtime"){
			fmt.Println("runtime error occurred")
		}
 		fmt.Println(str)
	}()
	//_,err := fmt.Println(os.Args[1])
	//if err != nil {
	//	panic(err)
	//}
	file ,err := os.Open("hello.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	fmt.Printf("%s",file)


}
