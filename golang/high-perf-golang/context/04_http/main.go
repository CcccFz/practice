package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func hanler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	select {
	case <- time.After(300 * time.Millisecond):
		fmt.Fprintf(w, "hello")
	case <- ctx.Done():
		err := ctx.Err()
		fmt.Println("err:", err)
	}
}

func server() {
	http.HandleFunc("/", hanler)
	http.ListenAndServe(":5678", nil)
}

func main() {
	go server()

	time.Sleep(500 * time.Millisecond)
	
	cli := http.Client{
		Timeout: 200 * time.Millisecond,
	}

	if rsp, err := cli.Get("http://localhost:5678"); err == nil {
		defer rsp.Body.Close()
		if data, err := io.ReadAll(rsp.Body); err == nil {
			fmt.Println(rsp.StatusCode, string(data))
		}
	} else {
		fmt.Println(err)
	} 
}