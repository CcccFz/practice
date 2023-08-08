package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/json-iterator/go"
)

type Book struct {
	Name   string `json:"title"`
	Price  float64
	Tags   []string
	Press  string
	Author People
}

type People struct {
	Name    string
	Age     int
	School  string
	Company string
	Title   string
}

var (
	people = People{
		Name:    "张三",
		Age:     18,
		School:  "清华大学",
		Company: "大乔乔教育",
		Title:   "开发工程师",
	}
	book = Book{
		Name:   "高性能golang",
		Price:  58.0,
		Tags:   []string{"golang", "编程", "计算机"},
		Press:  "机械工业出版社",
		Author: people,
	}
)

func TestStdJson(t *testing.T) {
	if data, err := json.Marshal(book); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(string(data))

		var book2 Book
		if err := json.Unmarshal(data, &book2); err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Printf("%+v\n", book2)
		}
	}
}

func TestJsoniterJson(t *testing.T) {
	if data, err := jsoniter.Marshal(book); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(string(data))

		var book2 Book
		if err := jsoniter.Unmarshal(data, &book2); err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Printf("%+v\n", book2)
		}
	}
}

func TestSonic(t *testing.T) {
	if data, err := sonic.Marshal(book); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(string(data))

		var book2 Book
		if err := sonic.Unmarshal(data, &book2); err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Printf("%+v\n", book2)
		}
	}
}

func BenchmarkStdJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(book)
		var book2 Book
		json.Unmarshal(data, &book2)
	}
}

func BenchmarkJsoniterJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := jsoniter.Marshal(book)
		var book2 Book
		jsoniter.Unmarshal(data, &book2)
	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := sonic.Marshal(book)
		var book2 Book
		sonic.Unmarshal(data, &book2)
	}
}

// go test -v ./golang/high-perf-golang/json -run=TestStdJson -count=1
// go test -v ./golang/high-perf-golang/json -run=JsoniterJson -count=1
// go test -v ./golang/high-perf-golang/json -run=TestSonic -count=1

// go test ./golang/high-perf-golang/json -bench=BenchmarkStdJson -run=none -count=1 -benchmem -benchtime=2s
// go test ./golang/high-perf-golang/json -bench=BenchmarkJsoniterJson -run=none -count=1 -benchmem -benchtime=2s
// go test ./golang/high-perf-golang/json -bench=BenchmarkSonic -run=none -count=1 -benchmem -benchtime=2s

//作者：高性能golang https://www.bilibili.com/read/cv25000934?spm_id_from=333.999.0.0 出处：bilibili