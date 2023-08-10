package main

import (
	"context"
	"fmt"
	"time"
)



func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100) //超时后会自动调用context的Deadline，Deadline会，触发Done
	defer cancel()
	select {
	case <- time.After(300 * time.Millisecond):
		fmt.Println("未超时")
	case <- ctx.Done(): //ctx.Done()是一个管道，context超时或者调用了cancel()都会关闭这个管道，然后读操作就会立即返回
		err := ctx.Err()        //如果发生Done（管道被关闭），Err返回Done的原因，可能是被Cancel了，也可能是超时了
		fmt.Println("超时:", err) //context deadline exceeded
	}
}