package main


type Stack []int

func (s *Stack) Push(v int){
	*s = append(*s, v)
}
func (s *Stack) Pop() int{
	v := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s)-1]
	return v
}
func (s *Stack) PrintStack(){
	for _,v := range *s{
		print(v," ")
	}
	println()
}
func (s *Stack) IsEmpty() bool{
	return (len(*s) == 0)
}
func main(){
	s := Stack{}
	s.Push(10)
	s.Push(104)
	s.PrintStack()
	println(s.Pop())
	println(s.Pop())
	println(s.IsEmpty())
}
