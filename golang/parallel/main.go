package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Pointer(new(interface{})))
	fmt.Println(unsafe.Pointer(new(interface{})))
}
