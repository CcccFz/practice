package _8_function

import "fmt"

// go函数不支持嵌套、重载、默认参数
// 支持无需声明原型、不定长参数、多返回值、命名返回值参数、匿名函数、闭包
// func 的 { 不能另起一行，func可作为类型， go语言一切皆类型

func Test() {
	ret1, ret2, ret3 := A(1, "hh")
	fmt.Println(ret1, ret2, ret3)

	B(1, 2, 3, 4)
	c := C
	c() // 一切皆类型

	// 匿名函数
	d := func() {
		fmt.Println("sdsdsfsf")
	}
	d()

	// 闭包
	f := closure(10)
	fmt.Println(f(5))
}

// 一个返回值，不需要括号
func A(a int, b string) (c, d int, e string) {
	c, e = a, b
	d = c
	return c, d, e
}

// 不定长参数, 必须是最后一个参数; 是个slice的值拷贝。改变不会影响
// 如果参数传递slice  func B(a []int)，就是个引用。改变会影响
func B(a ...int) {
	fmt.Println(a)
}

func C() {
	fmt.Println("HHHHHHH")
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
