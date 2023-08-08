package _13_reflection

import (
	"fmt"
	"reflect"
)

// 反射使得interface{}有更大的发挥余地
// 反射使用TypeOf和ValueOf函数从接口中获取目标对象信息，且必须为对象
// 反射将匿名字段作为独立字段(匿名字段本质)
// 想要通过反射修改对象状态，前提是interface.data是settable，即pointer-interface
// 通过反射可以“动态”调用方法

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello.")
}

func (u User) AA(name string) {
	fmt.Println("Hello", name, ". my name is", u.Name)
}

func Test() {
	u := User{1, "ok", 12}
	Info(u)

	m := Manager{User: User{2, "HH", 18}, title: "fongwsd"}
	t := reflect.TypeOf(m)
	fmt.Printf("\n%#v\n", t.FieldByIndex([]int{0, 2}))

	x := 123
	xv := reflect.ValueOf(&x)
	xv.Elem().SetInt(999)
	fmt.Println(x)

	Set(&u)
	fmt.Println(u)

	v := reflect.ValueOf(u)
	mm := v.MethodByName("AA")
	args := []reflect.Value{reflect.ValueOf("joe")}
	mm.Call(args)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Not Reflect")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v", f.Name, f.Type, val)
	}
	fmt.Println("\nMethods:")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Error type")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if f.IsValid() && f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
