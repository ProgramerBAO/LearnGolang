package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string
	Age      int
}

type Student struct {
	Class string
	// 这里是匿名表达，我们也可以写名字。
	User
}

func main() {
	s := Student{
		Class: "1",
		User: User{
			UserName: "Bob",
			Age:      18,
		},
	}

	check(s)
}

// 空接口
func check(inter interface{}) {
	// 拿到类型
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		// 按照层级取值
		fmt.Println(v.Field(i))
	}
	fmt.Println(t, v)
	// 按照层级取值  这里直接取第二层
	fmt.Println(v.FieldByIndex([]int{1}))
	// 根据name拿
	fmt.Println(v.FieldByName("Class"))
}
