package _5_array

import "fmt"

// 数组的长度也是数组的一部分，长度不同的数组，认为是不同类型
// 区分指向数组的指针和指针数组
// 不管是数组本身，还是指向数组的指针，都可以用下标存取值, 如p[0] = 10
// 数组在go中为值类型，传递给函数时，是值拷贝。想传递引用的话，用切片
// 数组可以用==和!=的比较，但是不能有>和<的比较

func Test() {
	var a [2]int
	b := [20]int{1}
	c := [20]int{19: 1} // [...]int{19: 1}
	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 指向数组的指针
	var p *[2]int = &a // p := new([2]int) 一样的效果
	fmt.Println(p)

	// 指针数组
	x, y := 1, 2
	e := [...]*int{&x, &y}
	fmt.Println(e)
}
