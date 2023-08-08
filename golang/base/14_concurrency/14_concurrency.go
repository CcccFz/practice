package _14_concurrency

import (
	"fmt"
	"runtime"
)

// goroutine 奉行通过通信来共享内存，而不是共享内存来通信
// channel是goroutine沟通的桥梁，大部分是阻塞同步的，是引用类型
// 可使用for range来迭代不断操作channel
// 可以设置单向或双向通道
// 可以设置缓存大小，在未被填满前不会发生阻塞

func Test() {
	c := make(chan bool) // 默认是双向通道
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	cc := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		<-cc
	}()
	cc <- true

	runtime.GOMAXPROCS(runtime.NumCPU())
	gc := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(gc, i)
	}
	for i := 0; i < 10; i++ {
		<-gc
	}
}

func Go(c chan bool, idx int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += 1
	}
	fmt.Println(idx, a)

	c <- true
}
