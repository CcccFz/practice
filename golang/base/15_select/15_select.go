package _15_select

import (
	"fmt"
)

// select
// 可处理一个或多个channel的发送与接收
// 同时有多个可用的channel时安随机顺序处理
// 可用空的select来阻塞main函数
// 可设置超时

func Test() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)

	for i := 0; i < 2; i++ {
		<-o
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
