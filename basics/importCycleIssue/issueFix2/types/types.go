package types

type Ageter interface {
	GetA() *A
}
type Cuser interface {
	UseC()
}
type Buser interface {
	UseB()
}
type A struct {
	B *B
	C *C
}
type B struct {
	O Ageter
	Cuser
}
type C struct {
	O Ageter
	Buser
}
