package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var cnt uint32
	wg := sync.WaitGroup{}

	for i := uint32(0); i < 10; i++ {
		wg.Add(1)

		go func(i uint32) {
			for {
				if atomic.LoadUint32(&cnt) == i {
					fmt.Println(i)
					atomic.AddUint32(&cnt, 1)

					wg.Done()
					break
				}

				time.Sleep(time.Nanosecond)
			}
		}(i)
	}

	wg.Wait()
}
