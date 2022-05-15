package Set

type Set interface {
	add(e int)
	remove(e int)
	contains(e int) bool
	getSize(e int) int
	isEmpty() bool
}
