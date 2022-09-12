package main

// 反射可以知道进来的数据类型的原始数据类型

import (
	"fmt"
	"reflect"
)

type User_reflect struct {
	Name   string
	Age    int
	Gender bool
}

type Student_reflect struct {
	Class string
	User_reflect
}

func (u User_reflect) SayName(name string) {
	fmt.Println("My name is", name)
}

func check_reflect(inter interface{}) {
	// 获取输入的参数接口中的数据的值
	v := reflect.ValueOf(inter)
	// 获取输入的参数接口中的数据的值的类型
	t := reflect.TypeOf(inter)
	fmt.Println(v, t)
	// 获取结构体中的第几个属性
	p := v.Elem()
	if p.CanSet() {
		for i := 0; i < p.NumField(); i++ {
			// 因为传入的是指针，所以要用Elem() 先把数据取出来
			fmt.Println(p.Field(i))
		}
	}
	ty := t.Kind()
	if ty == reflect.Struct {
		fmt.Println("我是 Struct")
	}
	if t.String() == "*main.Student_reflect" {
		// 接收原始数据
		e := v.Elem()
		e.FieldByName("Class").SetString("202")
		fmt.Println(inter)
	}
	// 获取方法
	if p.Type().String() == "main.User_reflect" {
		m := p.Method(0)
		m.Call([]reflect.Value{reflect.ValueOf("Alic")})
	}

}

func main() {
	u_reflect := User_reflect{
		Name:   "Bob",
		Age:    20,
		Gender: true,
	}
	check_reflect(&u_reflect)

	s_reflect := Student_reflect{
		Class: "102",
		User_reflect: User_reflect{
			Name:   "Kenny",
			Age:    20,
			Gender: true,
		},
	}
	check_reflect(&s_reflect)
	fmt.Println(s_reflect)
}
