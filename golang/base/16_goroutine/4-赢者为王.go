package main

import "fmt"

func searchByGoogle(s string) string {
	return s + "in google"
}

func searchByBing(s string) string {
	return s + "in bing"
}

func searchByBaidu(s string) string {
	return s + "in baidu"
}

func main() {
	ch := make(chan string, 10)
	key := "golang"

	go func() {
		ch <- searchByGoogle(key)
	}()

	go func() {
		ch <- searchByBing(key)
	}()

	go func() {
		ch <- searchByBaidu(key)
	}()

	fmt.Println(<-ch)
}
