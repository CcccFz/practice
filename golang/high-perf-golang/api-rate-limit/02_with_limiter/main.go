package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

var (
	concurrent int32
)

func hadler() {
	atomic.AddInt32(&concurrent, 1)
	time.Sleep(50 * time.Millisecond)
	return
}

func callHandler() {
	limiter := rate.NewLimiter(rate.Every(100 * time.Millisecond), 1)
	for {
		limiter.WaitN(context.Background(), 1)
		hadler()
	}
	// for {
	// 	time.Sleep(limiter.ReserveN(time.Now(), 1).Delay())
	// 	hadler()
	// }
	// for {
	// 	if limiter.AllowN(time.Now(), 1) {
	// 		hadler()
	// 	}
	// }
}

func main() {
	go callHandler()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println("过去一秒调用接口次数：", atomic.LoadInt32(&concurrent))
		atomic.StoreInt32(&concurrent, 0)
	}
}
