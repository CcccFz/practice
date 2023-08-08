package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func example1()  {
	var mu sync.Mutex
	mu.Lock()

	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()

	mu.Lock() // 此处没有Lock，不会打印
}

func example2() {
	ch := make(chan bool)

	go func() {
		fmt.Println("hello world")
		<- ch
	}()

	ch <- true

	// 此种方式，若ch改为缓冲队列，不会打印
}

func example3() {
	ch := make(chan bool)

	go func() {
		fmt.Println("hello world")
		ch <- true
	}()

	<- ch
}

// range
func example4() {
	ch := make(chan int, 10)

	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			fmt.Println("hello world: " + strconv.Itoa(i) )
			ch <- i
		}(i)
	}

	// range报错是看有无协程还在，不是看有无数据。即使缓冲chan中装满数据，但所有协程都结束了，range仍要报错。
	for i := range ch {
		fmt.Println(i)
	}
}

// close
func example5() {
	ch := make(chan int, 10)

	go func() {
		defer close(ch)

		x, y := 1, 1
		for i := 0; i < cap(ch); i++ {
			ch <- x
			x, y = y, x + y
		}
	}()

	// 如果没有close，会阻塞在此处，并且协程会立即完完毕，报all goroutines are asleep - deadlock!
	// 有close会把数据读完
	for i := range ch {
		fmt.Printf("%d ", i)
	}
}

func example6() {
	ch := make(chan int, 10)

	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			fmt.Println("hello world: " + strconv.Itoa(i) )
			ch <- i
		}(i)
	}

	// 等待所有协程完成
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<- ch)
	}
}

func example7() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println("hello world: " + strconv.Itoa(i))
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()
}

func example8() {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			runtime.Gosched()   // runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
			fmt.Println(s)
		}
	}

	go say("world")
	say("hello")
}

func example9()  {
	ch := make(chan int)
	a := []int{7, 2, 8, -9, 4, 0}

	sum := func(a []int) {
		total := 0
		for _, v := range a {
			total += v
		}
		fmt.Println(total)
		ch <- total
	}

	go sum(a[:len(a)/2])
	go sum(a[len(a)/2:])

	x, y := <- ch, <- ch
	fmt.Printf("%d + %d = %d", x, y, x + y)
}

func example10() {
	ch := make(chan int, 1)  // buffer大小 小于 写入的次数，会报 fatal error: all goroutines are asleep - deadlock!
	ch <- 1
	ch <- 2     // 到此处时，chan满，发送数据阻塞改main协程，不会再往下执行；但发现，已经没有其它协程，不可能有协程接收数据，则会一直阻塞。
	            // 程序已知道，会死锁，且不可能有人来救场，觉得很尴尬，不如就直接杀掉程序。报错，所有协程都睡了
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}

// 超时
func example11() {
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <- time.After(3 * time.Second):
				fmt.Println("timeout")
				ch <- true
			}
		}
	}()
	<- ch
}

// 多个channel来控制关闭
func example12() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<- ch)
		}
		quit <- true
	}()

	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x + y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func example13() {
	ch := make(chan bool)

	go func() {
		for {
			select {
			case <- ch:
			default:         // 如果没有default的话，select会被阻塞，不会打印get it
			}
			fmt.Println("get it")
		}
	}()

	for {
		_ = ""
	}
}

func example14() {
	ch := make(chan int, 10)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case ch <- 1:     // 随机的
			case ch <- 2:     // 随机的
			case <- quit:
				fmt.Println("quit")
				return
			}
		}
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	time.Sleep(3 * time.Second)
	quit <- true
	time.Sleep(time.Second)
}

// close后，再读取channel。先读完剩下的值，之后读出为0值
func example15() {
	ch := make(chan bool)
	close(ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println()

	a := make(chan int, 10)
	a <- 1
	a <- 2
	close(a)
	fmt.Println(<- a)
	fmt.Println()

	v, ok := <- a
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println()

	v, ok = <- a
	fmt.Println(v)
	fmt.Println(ok)
}

func example16() {
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer fmt.Println("quit" + strconv.Itoa(i))

			for {
				select {
				default:
				case <- ch:
					return
				}
			}
		}(i)
	}

	time.Sleep(time.Second)
	close(ch)   // close时会向ch广播发命令。但如果协程中涉及资源清理，清理会来不及，main会立即结束
	// 需要保证
}

// runtime goroutine
// Goexit: 退出当前执行的goroutine，但是defer函数还会继续调用
// Gosched: 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
// NumCPU: 返回CPU核数量
// NumGoroutine: 返回正在执行和排队的任务总数
// GOMAXPROCS: 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	//example5()
	//example6()
	//example7()
	//example8()
	//example9()
	//example10()
	//example11()
	//example12()
	//example13()
	//example14()
	//example15()
	example16()
}
