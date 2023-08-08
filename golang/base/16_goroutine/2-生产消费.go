package main

import (
	"fmt"
	"time"
)

// 并发模式

func Producer(ch chan int, n int) {
	for i := 0;; i++ {
		ch <- i * n
	}
}

func Consumer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main()  {
	ch := make(chan int, 10)

	go Producer(ch, 2)
	go Producer(ch, 5)
	go Consumer(ch)

	time.Sleep(2 * time.Second)
}