package main

import (
	"fmt"
	"time"
)

//const PI = 3.14

// 导入多个包，声明多个常量、全局变量、一般类型（非接口、非结构）都可以使用这种形式；
// 这种形式叫组：全局变量组(只有全局变量组，没有局部变量组)、常量组、一般类型组
// 全局变量不能省略var
var (
	name = "gopher"
	num  = 1
)

//type intType int

//type gopher struct {
//
//}

//type golang interface {
//
//}

// 首字母大写，为public
func main() {
	fmt.Println("HHH")
	fmt.Println(time.Now())
	println(num)
	println(name)
}
