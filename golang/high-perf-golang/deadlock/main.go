package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	ch := make(chan struct{}, 0)
	go func ()  {
		ch <- struct{}{}
	}()
	ch <- struct{}{}
}

// 不知道sleep有什么代码，可能有希望不足赛
func main11() {
	ch := make(chan struct{}, 0)
	go func ()  {
		ch <- struct{}{}
	}()
	time.Sleep(time.Minute)
	ch <- struct{}{}
}

func main2() {
	ch := make(chan struct{}, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		ch <- struct{}{}
	}()
	wg.Wait()
}

// 不知道sleep有什么代码，可能有希望不足赛
func main22() {
	ch := make(chan struct{}, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		ch <- struct{}{}
	}()
	go func() {
		time.Sleep(time.Minute)
	}()
	wg.Wait()
}

// 无缓冲chan的读写，必须出现在2个不同的go
func main3() {
	ch := make(chan struct{}, 0)
	ch <- struct{}{}
	fmt.Println("")
}

func main() {
	// main1()
	// main11()
	// main2()
	// main22()
	main3()
}