package _2_type

import (
	"math"
	"strconv"
)

// int/uint 根据平台来确定是32位还是64位
// int8/uint8 范围为-128~127/0~255
// 字节型byte 是uint8的别名
// rune是int32的别名，常用语unicode字符串

// 浮点型 float32/float64，4/8字节，精确到7/15小数位，没有double型
// 复数型 complex64/complex128，8/16字节
// 其它类型 array、struct、string
// 引用类型 slice、map、chan
// 接口类型 interface
// 函数类型 func
// uintptr 本质是无符号整型，用于保存指针地址

func Test() {
	// 默认值
	var a int
	var b string
	var c bool
	var d [1]int // 真正数组要有数组大小，否则会变为切片
	println(a)
	println(b)
	println(c)
	println(d[0])

	// 检查类型的最小值，最大值
	println(math.MaxInt16)

	// :就是var的意思，如果是之前有的变量不用:号
	_, bb, cc := 1, 2, "3"
	println(bb, cc)

	// go不存在隐式转换，且必须兼容转换（比如bool转int不兼容）
	dd := 100.1
	ddd := int(dd)
	println(ddd)

	e := 65
	println(string(e))
	println(strconv.Itoa(e))
}
