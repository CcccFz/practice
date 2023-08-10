package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

var (
	text = "recall 74 by radic, recall 0 by milvus_short, recall 0 by milvus_long\n"
)

type BufferedFileWriter struct {
	f           *os.File
	cache       []byte
	cacheEndIdx int
}

func NewBufferedFileWriter(f *os.File, size int) *BufferedFileWriter {
	return &BufferedFileWriter{
		f:           f,
		cache:       make([]byte, size),
		cacheEndIdx: 0,
	}
}

func (w *BufferedFileWriter) Write(data []byte) {
	if len(data) >= len(w.cache) {
		w.Flush()
		w.Write(data)
		return
	}
	if w.cacheEndIdx+len(data) > len(w.cache) {
		w.Flush()
	}
	copy(w.cache[w.cacheEndIdx:], data)
	w.cacheEndIdx += len(data)
}

func (w *BufferedFileWriter) Flush() {
	if w.cacheEndIdx == 0 {
		return
	}
	w.f.Write(w.cache[:w.cacheEndIdx])
	w.cacheEndIdx = 0
}

func WriteFile(path string, isBuffer bool) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if isBuffer {
		w := NewBufferedFileWriter(f, 4096)
		defer w.Flush()
		for i := 0; i < 100_000; i++ {
			w.Write([]byte(text))
		}
	} else {
		for i := 0; i < 100_000; i++ {
			f.Write([]byte(text))
		}
	}
}

func OneWriteFile(path string, data []byte) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(data)
}

func genBigBytes() []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < 100_000; i++ {
		buf.Write([]byte(text))
	}
	return buf.Bytes()
}


func main() {
	start := time.Now()
	WriteFile("a.txt", false)
	fmt.Printf("直接写入耗时：%dms\n", time.Since(start).Milliseconds())

	start = time.Now()
	WriteFile("b.txt", true)
	fmt.Printf("有缓冲写入耗时：%dms\n", time.Since(start).Milliseconds())

	start = time.Now()
	OneWriteFile("c.txt", genBigBytes())
	fmt.Printf("一次性写入耗时：%dms\n", time.Since(start).Milliseconds())
}
