package main

import (
	"sync"
)

const (
	minSize = 10
	maxSize = 163840
	factor  = 2
)

var intSlicePool *SlicePool

type SlicePool struct {
	pools   []sync.Pool
	sizes   []int
	minSize int
	maxSize int
}

func init() {
	n := 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		n++
	}

	intSlicePool = &SlicePool{
		pools:   make([]sync.Pool, n),
		sizes:   make([]int, n),
		minSize: minSize, maxSize: maxSize,
	}

	n = 0
	for chunkSize := minSize; chunkSize <= maxSize; chunkSize *= factor {
		intSlicePool.sizes[n] = chunkSize
		intSlicePool.pools[n].New = func(size int) func() interface{} {
			return func() interface{} {
				buf := make([]int, 0, size)
				return &buf
			}
		}(chunkSize)
		n++
	}
}

func Alloc(size int) []int {
	if size > maxSize {
		return make([]int, 0, size)
	}

	i := 0
	for ; i < len(intSlicePool.sizes); i++ {
		if size <= intSlicePool.sizes[i] {
			break
		}
	}

	return (*intSlicePool.pools[i].Get().(*[]int))[:0]
}

func Free(s []int) {
	size := cap(s)
	if size > intSlicePool.maxSize {
		return
	}

	i := 0
	for ; i < len(intSlicePool.sizes); i++ {
		if size <= intSlicePool.sizes[i] {
			break
		}
	}

	intSlicePool.pools[i].Put(&s)
}

func main() {
	s := Alloc(27)

	for i := 0; i < 27; i++ {
		s = append(s, i)
	}

	defer Free(s)
}
