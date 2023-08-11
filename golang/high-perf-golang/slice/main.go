package main

import "fmt"

func f1() {
	a1 := make([]int, 0, 10)
	b1 := a1
	fmt.Printf("a1: %p\n", &a1)
	fmt.Printf("b1: %p\n", &b1)	
}

func main() {
	f1()
}