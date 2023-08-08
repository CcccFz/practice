package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type once struct {
	do   int32
	lock sync.Mutex
}

func (o *once) run(f func()) {
	if atomic.LoadInt32(&o.do) == 1 {
		return
	}

	o.lock.Lock()
	defer o.lock.Unlock()
	if o.do == 0 {
		defer atomic.StoreInt32(&o.do, 1)
		f()
	}
}

func main() {
	var o once
	var wg sync.WaitGroup
	t := time.Now().UnixNano()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			o.run(printf)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("耗时: %d\n", time.Now().UnixNano()-t)
}

func printf() {
	fmt.Println(1)
}
