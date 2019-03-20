package i

import (
)

type A interface {
	GetB() B
	GetC() C
}

type B interface {}

type C interface {}
