package _12_interface

import "fmt"

// 接口是一个或多个方法签名的集合
// 只要实现了接口的所有方法签名，就算实现接口。无需声明实现了哪个接口
// 接口可以匿名嵌入其它接口，或嵌入到结构中
// 将对象赋值给接口时，会发生拷贝，而接口内部存储是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
// 只有当接口存储的类型和对象都为nil时，接口才等于nil
// 接口调用不会做receiver的自动转换，p.func()和(*p).func()必须区分
// 接口同样支持匿名字段方法
// 接口可实现类似OOP中的多态，go没有继承
// 空接口可作为任何类型数据的容器，相当于是根类型

// 类型断言
// 通过类型断言的ok pattern可以判断接口中的数据类型
// 使用type switch则可针对空接口进行比较全面的类型判断

// 接口转换
// 可以将拥有超集的接口转换为子集的接口

type USB interface {
	Name() string
	Connector // 嵌入接口
}

type Connector interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect", pc.name)
}

func Test() {
	u := PhoneConnecter{"PhoneConnecter"}
	u.name = "PhoneConnecter1111111111"
	u.Connect()
	Disconnect(u)
	Disconnect2(u)

	p := PhoneConnecter{"PhoneConnecter"}
	// 将对象赋值给接口时，会发生拷贝，而接口内部存储是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
	a := Connector(p)
	p.name = "PhoneConnecter22222222222222222222222"
	a.Connect()

	var i1 interface{}
	fmt.Println(i1 == nil)

	var q *int = nil
	i1 = q
	fmt.Println(i1 == nil)
}

func Disconnect(usb USB) {
	// 类型判断语法糖
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect.", pc.name)
		return
	}
	fmt.Println("Unknown device.")
}

// 可以func Disconnect(usb interface{}) 任何类型都可以传入空接口
func Disconnect2(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect.", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}
