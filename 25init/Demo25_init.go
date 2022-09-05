package demo25_init

import "fmt"

/*
	init函数没有参数，没有返回值 一个包中包含多个init时调用顺序是不确定的
	不允许手动调用
	在导入包的时候 如果指向使用他的init函数 可以在导包之前加入一个下划线
	_ "包的地址"  就类似于这样
*/
func init() {
	fmt.Println("first init")
}

func init() {
	fmt.Println("second init")
}

func Sub(a, b int) int {
	return a - b
}
