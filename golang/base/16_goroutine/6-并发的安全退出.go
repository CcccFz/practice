package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// channel版本
func quit1() {
	var wg sync.WaitGroup
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer fmt.Println("quit: " + strconv.Itoa(i))

			for {
				select {
				default:
				case <- ch:
					return
				}
			}
		}(i)
	}

	close(ch)
	wg.Wait()
}

// context版本
func quit2() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer fmt.Println("quit: " + strconv.Itoa(i))

			for {
				select {
				default:
				case <- ctx.Done():
					return
				}
			}
		}(i)
	}

	cancel()
	wg.Wait()
}

func main() {
	quit1()
	quit2()
}