package _6_stack

type Node struct {
	next *Node
	val interface{}
}

type LinkedListStack struct {
	top *Node
}
