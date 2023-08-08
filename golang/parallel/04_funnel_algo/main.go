package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var (
	pool     = make([][]byte, 20)
	poolSize = len(pool)
	cnt      int
	m        runtime.MemStats
)

type LeakyBucket struct {
	size       int
	bufferChan chan []byte
}

func NewLeakyBucket(size int) *LeakyBucket {
	return &LeakyBucket{size: size, bufferChan: make(chan []byte, 100)}
}

func (bucket *LeakyBucket) Get() (b []byte) {
	select {
	case b = <-bucket.bufferChan:
		return b
	default:
		cnt++
		return make([]byte, bucket.size)
	}
}

func (bucket *LeakyBucket) Push(b []byte) {
	select {
	case bucket.bufferChan <- b:
	default:
	}
}

func main() {
	//for ;; {
	//	i := rand.Intn(poolSize)
	//	pool[i] = make([]byte, 1e7)
	//	printf()
	//}

	buffer := NewLeakyBucket(1e7)
	for {
		b := buffer.Get()

		i := rand.Intn(poolSize)
		if pool[i] != nil {
			buffer.Push(pool[i])
		}

		pool[i] = b
		printf()
	}
}

func printf() {
	time.Sleep(100 * time.Millisecond)

	total := 0
	for i := 0; i < poolSize; i++ {
		if pool[i] == nil {
			continue
		}
		total += len(pool[i])
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("HeapSys:%d,bytes:%d,HeapAlloc:%d,HeapIdle:%d,HeapReleased:%d,cnt:%d\n",
		m.HeapSys/1e6, total/1e6, m.HeapAlloc/1e6, m.HeapIdle/1e6, m.HeapReleased/1e6, cnt)
}
