package _6_slice

import "fmt"

// 指向底层数组，变长数组的替代方案。可以关联部分或全部数组
// len()取元素个数, cap()取容量
// 一般使用make创建， make([]T, len, cap)，cap省略表示和len相同
// 多个slice指向一个数组，其中一个改变，会影响到所有

func Test() {
	var s1 []int // 空slice
	fmt.Println(s1)

	a := [10]int{}
	fmt.Println(a)
	s2 := a[5:]
	fmt.Println(s2)
	s3 := make([]int, 3, 10) // 10个内存块，初始化3个，超过10个时，指向一个20个内存块的新数组(2倍)
	fmt.Println(s3, len(s3), cap(s3))

	// 指向切片的切片 reslice
	// 索引不可超过 基slice的最大容量，不会重新分配底层数组
	b := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	sb1 := b[2:5]
	sb2 := sb1[1:3]
	sb3 := sb1[3:5]
	fmt.Println(sb1)
	fmt.Println(sb2, len(sb2), cap(sb2))
	fmt.Println(sb3, len(sb3), cap(sb3))

	sb4 := sb1[1:]
	fmt.Println(sb4)

	// append，拼接后的超过slice的容量则会重新分配数组
	c := make([]int, 3, 6)
	fmt.Printf("%p\n", c)
	c = append(c, 1, 2, 3, 4) // 区别于c = append(c, 1, 2, 3)
	fmt.Printf("%v %p\n", c, c)

	// 共享数组，一个slice改变就会导致另一个slice变
	d := []int{1, 2, 3, 4, 5}
	sd1 := d[1:3]
	sd2 := d[2:5]
	// 但是如果一个slice通过append超过容量，slice就指向新的数组，所以二者就没共享同一个数组了
	sd2 = append(sd2, 6)
	sd2[0] = 9
	fmt.Println(sd1, sd2)

	// copy
	e1 := []int{1, 2, 3, 4, 5, 6}
	e2 := []int{7, 8, 9}
	copy(e1, e2) // 第一个是des, 第二个是src, 改变的是des。
	//copy(e1[1:2], e2[:2])  // des和src都可以用切片的方式，指定考src的哪部分，到desc的哪部分
	// copy不会改变数组，只是值得替换
	fmt.Println(e1)
}
