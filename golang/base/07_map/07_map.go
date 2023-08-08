package _7_map

import (
	"fmt"
	"sort"
)

// Key必须是支持==或!=的类型，不可以是函数、map、slice
// map查找比线性块，但比索引访问慢100倍
// map使用make创建，支持:=  make(map[keyType]valueType, cap)

func Test() {
	var a map[int]string
	a = map[int]string{} // make(map[int]string)
	b := map[int]string{}
	fmt.Println(a, b)

	a[1] = "ok"
	fmt.Println(a, a[1], a[2]) // 没有的key就是空

	var c map[int]map[int]string = make(map[int]map[int]string)
	c[1] = map[int]string{}
	c[1][1] = "hh"
	fmt.Println(c, c[2][1])

	_, ok := c[2][1]
	if !ok {
		c[2] = map[int]string{}
	}
	c[2][1] = "JJ"
	fmt.Println(c)

	// 迭代操作
	sm := make([]map[int]string, 5)
	// 不能用for _, v := range sm {}, 因为v是个拷贝值
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
	}
	fmt.Println(sm)

	// sort
	um := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(um))
	i := 0
	for k := range um {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

	w := make(map[string]int)
	for k, v := range um {
		w[v] = k
	}
	fmt.Println(w)
}
