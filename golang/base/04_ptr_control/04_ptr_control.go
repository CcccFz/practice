package _4_ptr_control

// 指针
// Go指针不支持->，替代的是"."
// ++和--只能作为语句放在单独一行，且必须在变量右边；不能作为表达式

// if
// 条件表达式不能有括号
// 支持初始化表达式，但是作用域只在if内

// 循环 只有for
// for {} 无限循环
// for a <= 3 {} 条件循环
// for i := 0; i < 3; i++ {} 经典循环

// switch 选择
// 不写break，选中后默认终止，若是想继续需要fallthrough语句

// 跳转 goto, break, continue
// 三个语法都可以配合标签使用，标签大小写敏感，不用会编译错误
// break与continue配合标签可以跳出多层循环
// goto是调整执行位置，与其他两个配合标签的效果不一样

func Test() {
	name := 1
	var p *int = &name
	println(p)

	// a只在if中有效
	if a := name; a > 0 {
		println("a > name")
	}

	for i := 0; i < 3; i++ {
		println(i)
	}

	// switch
	b := name - 1
	switch {
	case b == 0:
		println("b == 0")
		b = 1
		fallthrough
	case b >= 0:
		println("b: ", b)
	default:
		println("No match")
	}

LABEL1:
	// breake
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1 // 如果这里是goto，那么还是死循环；goto时，label放在循环后面
			}
		}
	}

LABEL2:
	// continue
	for i := 0; i < 10; i++ {
		for {
			continue LABEL2
			println("after")
		}
	}
	println("OK")
}
