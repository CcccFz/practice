package _3_const_operation

// 常量是在编译时值就已经确定的
// 等号右侧必须是常量或者常量表达式
// 常量表达式中的函数必须是内置函数

// 常量的初始化规则与枚举
// 在常量组中，如果不提供初始值，则表示使用上一个表达式的值
// 使用相同的表达式，不代表具有相同的值
// iota是常量计数器，从0开始，组中每定义1个常量自动递增（可以用来实现枚举）
// 每遇到一个const关键字，iota就会重置为0

//const a int = 1
//const b = 'A'

//const (
//	c = a
//	d = a + 1
//	e, f = b + 1, b + 2
//)

//const (
//	g = 1
//	h
//	i
//)

// 错误，每一行必须一致
//const (
//	a, b = 1, 2
//	h
//)

const (
	a = 'A'
	b
	c = iota
	d
)

// 计算机存储单位
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func Test() {
	println(a, b, c, d)
	println(KB)
	println(MB)
	println(GB)
	println(TB)
}
