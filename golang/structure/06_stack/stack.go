package _6_stack

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Top() interface{}
	Flush()
}
