package main

import (
	"fmt"
	"time"
)


func hadler() string {
	time.Sleep(200 * time.Millisecond)
	return "ok"
}

func main() {
	var (
		rsp string
		done = make(chan struct{}, 1)
	)

	go func () {
		rsp = hadler()
		done <- struct{}{}
	}()

	select {
	case <- done:
	case <- time.After(100 * time.Millisecond):
		rsp = "timeout"
	}

	fmt.Println(rsp)
}