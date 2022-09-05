package main

import (
	"fmt"
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

}
