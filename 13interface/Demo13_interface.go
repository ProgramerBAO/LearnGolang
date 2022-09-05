package main

import "fmt"

type Study interface {
	learn()
}

type Student struct {
	name string
	age  int
}

// 方法实现接口
func (s Student) learn() {
	fmt.Println(s.name + " good good study, day day up")
}

func main() {
	/*
		type interfaceName interface {
			method()
			method()
		}
	*/
	stu := Student{
		name: "Bob",
		age:  18,
	}
	stu.learn()
}
