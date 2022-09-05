package main

import "fmt"

func main() {
	// 静态类型 变量声明时候的类型
	var number int = 20
	fmt.Println(number)

	// 动态类型 程序运行时才能看得见
	var in interface{}
	in = 100
	fmt.Println(in)
	in = "go"
	fmt.Println(in)
}
