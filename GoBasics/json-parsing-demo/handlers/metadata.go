package handlers

type JSONData struct {
	Student []Student	`json:"student"`
}
type Student struct {
	Name string	`json:"Name"`
	Id   string	`json:"Id"`
	Age  int	`json:"Age"`
	Marksheet Marksheet	`json:"Marksheet"`
}

type Marksheet struct{
	Maths	int	`json:"Maths"`
	Physics	int	`json:"Physics"`
	Chemistry	int	`json:"Chemistry"`
	Biology	int	`json:"Biology"`
}