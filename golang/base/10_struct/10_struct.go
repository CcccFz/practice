package _10_struct

import "fmt"

// Go没有class
// type <Name> struct{}
// 支持指向自身指针类型成员
// 支持匿名结构，可用作成员变量或定义成员变量，匿名结构也可以用于map的值
// 相同类型的成员可进行直接拷贝赋值
// 只支持==或!=的比较，但限于相同类型
// 支持匿名字段，本质上是定义以某个类型名为名称的字段，但是顺序不能错
// 嵌入结构作为匿名字段看起来像继承，但不是继承
// 可以使用匿名字段指针

type human struct {
	Sex int
}

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
	human
}

func Test() {
	a := person{Name: "CcccFz", human: human{Sex: 1}} //  这样写是为了解决初始化时名字冲突
	a.Age = 18
	a.Sex = 2 // a.human.Sex = 3也可以；a.Sex先找外层，找不到的话到嵌入的字段去找
	a.Contact.Phone = "135688"
	A(a) // 通过指针可以传递引用，或者初始化时用*person（更习惯用这种）
	fmt.Println(a)

	b := &struct {
		Name string
		Age  int
	}{
		Name: "Siyi", Age: 25,
	}
	fmt.Println(b)
}

func A(per person) {
	per.Age = 13
	fmt.Println("A", per)
}
