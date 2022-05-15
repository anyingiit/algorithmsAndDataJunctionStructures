package EInt

import "strconv"

type EInt struct {
	data int
}

func NewE(e int) *EInt {
	return &EInt{data: e}
}

func (e *EInt) String() string {
	return strconv.Itoa(e.data)
}

func (e *EInt) Id() int {
	return e.data
}

func (e *EInt) Data() interface{} {
	return e.data
}
