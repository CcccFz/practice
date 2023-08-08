package _2_regexp

import (
	"fmt"
	"regexp"
	"runtime"
)

func Test() {
	r, err := regexp.Compile(`Hello`)
	if err != nil {
		_, _, line, _ := runtime.Caller(0)
		fmt.Printf("There is a problem with your regexp in line: %d.\n", line)
		return
	}

	fmt.Println("相当于python的search，而不是python的match")
	if r.MatchString("Regular Hello Expression.") == true {
		fmt.Println("Match ")
	} else {
		fmt.Println("No match ")
	}

	fmt.Println("输出第一个找到匹配到的，没有匹配到则返回空")
	fmt.Println(r.FindString("Regular Hello Expression. Hullo again."))

	fmt.Println("验证输出第一个找到匹配到的")
	r, err = regexp.Compile(`.at`)
	fmt.Println(r.FindString("The cat sat on the mat."))

	fmt.Println("输出所有匹配到的，-1代表匹配直到末尾")
	fmt.Println(r.FindAllString("The cat sat on the mat.", -1))

	fmt.Println("匹配转义，匹配C:\\")
	fmt.Println(regexp.Compile("C:\\\\"))

	fmt.Println("得到匹配所在的位置")
	r, err = regexp.Compile(`well$`)
	fmt.Println(r.FindStringIndex("All is well that ends well"))
	r, err = regexp.Compile(`well`)
	fmt.Println(r.FindAllStringIndex("All is well that ends well", -1))

	fmt.Println("匹配多个表达式||的使用")
	r, err = regexp.Compile(`Jim|Tim`)
	fmt.Println(r.MatchString("Dickie, Tom and Tim"))
	fmt.Println(r.MatchString("Jimmy, John and Jim"))

	s := "Clara was from Santa Barbara and Barbara was from Santa Clara"
	// 两个正则
	r, err = regexp.Compile(`Santa Clara|Santa Barbara`)
	fmt.Println(r.FindAllStringIndex(s, -1))

	// 与上面效果一样，一个正则，两个单词
	r, err = regexp.Compile(`Santa (Clara|Barbara)`)
	fmt.Println(r.FindAllStringIndex(s, -1))
}
