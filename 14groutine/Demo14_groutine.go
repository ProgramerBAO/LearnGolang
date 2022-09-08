package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintInfo() {
	fmt.Println("go")
}

func PrintNumber(num int) {
	for i := 0; i < 3; i++ {
		fmt.Println(num)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	/*
		定义函数
		function functionName(parameterList)	{
			业务逻辑代码
		}
		开启函数
		functionNam(parameterList)
		开启一个协程调用函数
		go functionNam(parameterList)
	*/

	// go PrintInfo()
	go PrintNumber(1)

	go PrintNumber(2)

	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("hello world")
	time.Sleep(time.Second)

	// 使用协程管理器
	var wg sync.WaitGroup
	// 告诉它有多少协程
	wg.Add(1)
	// 传入地址 这样的话可以在方法内部改变外部的数据
	go Run(&wg)
	wg.Wait()
}

func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}
