package E

import "fmt"

type E interface {
	fmt.Stringer
	Id() int
	Data() interface{}
}
