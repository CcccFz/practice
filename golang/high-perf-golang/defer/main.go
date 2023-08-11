package main

import "fmt"

func f1() {
	a := "start"
	defer func() {
		fmt.Println(a)
	}()
	defer fmt.Println(a)
	a = "end"
}

func main() {
	f1()
}
