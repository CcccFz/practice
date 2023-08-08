package _11_method

import "fmt"

// go中没有class，但是依旧有method
// 通过显示receiver来实现与某个类型的组合
// 只能为同一个包中的类型定义方法
// receiver可以是值，也可以是指针
// 不存在方法的重载
// 可以使用值和指针来调用方法，编译过程自动完成转换
// 从某种意义上说，方法是函数的语法糖，因为receiver其实就是方法所接受的第一个参数
// Method Value vs Method Expression
// 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
// 类型别名不会拥有底层类型所附带的方法
// 方法可以调用结构中的公开字段

type A struct {
}

type B struct {
}

type TZ int

func Test() {
	a := A{}
	a.p()  // Method Value
	A.p(a) // Method Expression

	var b TZ
	b.Increase(100)
	fmt.Println(b)
}

// 不能func (a A) p(a int){}，因为go不能重载
// func (a *A) p()也可以，是引用传递，可以修改字段
func (a A) p() {
	fmt.Println("A")
}

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}
