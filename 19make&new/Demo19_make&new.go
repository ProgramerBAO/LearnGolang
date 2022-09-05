package main

import "fmt"

/*
	new 为所有类型分配内存，并初始化0值，返回指针
	make 只能为 slice，map，channel分配内存，并且初始化，返回的是类型
*/
func main() {
	num := new(int)
	fmt.Println(*num)
	// make只能初始化切片 map 和channel的数据 返回类型与其参数类型相同
	a := make([]int, 2, 10) // 切片的长度 切片的容量 容量不能小于长度
	b := make(map[string]int)
	c := make(chan int, 10)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
