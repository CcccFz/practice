package _9_defer

import "fmt"

// defer类似析构函数，在函数执行结束时，按调用顺序的相反顺序执行
// 函数发生错误时，也会执行defer
// 支持匿名函数中调用，常用于资源清理，文件关闭，解锁以及记录时间
// 通过与匿名函数配合可在return之后修改函数计算结果
// 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时已经获得了拷贝，否则是引用某个变量的地址

// Go 没有异常机制，但有panic/recover模式来处理错误
// Panic 可在任何地方引发，但recover只有在defer调用的函数中有效

func Test() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)

		defer func() {
			fmt.Println(i)
			// 这里形成了个闭包，当函数结束时，i为三，所以三次都打出3
		}()
		// 整个函数结束，才会逆向执行defer
	}

	A()
	B()
	C()

	fs := [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}

func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
