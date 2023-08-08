package main

import (
	"fmt"
	"time"
)

func handler(x, y int) string {
	now := time.Now()
	defer func ()  {
		fmt.Println(time.Since(now).Milliseconds())
	}()

	if x > y {
		time.Sleep(100 * time.Millisecond)
		return ""
	} else {
		time.Sleep(200 * time.Millisecond)
		return ""
	}
}

func main() {
	handler(3, 4)
}