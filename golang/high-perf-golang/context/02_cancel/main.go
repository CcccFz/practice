package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel() //调用cancel，关闭chan触发Done
	}()
	
	select {
	case <- time.After(300 * time.Millisecond):
		fmt.Println("超时")
	case <- ctx.Done():         //ctx.Done()是一个chan，调用了cancel()都会关闭这个chan，然后读操作就会立即返回
		err := ctx.Err()        //如果发生Done（chan被关闭），Err返回Done的原因，可能是被Cancel了，也可能是超时了
		fmt.Println("未超时:", err) //context canceled
	}
}