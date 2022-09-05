package main

import "fmt"

/*
	func functionName(parameterList)(resultList){
		// 函数体
		参数列表和返回值列表不是必须的
	}

	// 匿名函数
	func (parameterList)(resultList){
		// 函数体
		参数列表和返回值列表不是必须的
	}

	函数可见性
	1 函数首字母大写，对于所有的包都是public的，其他包可以任意调用
	2 函数首字母是小写的，这个函数是private的，其他包无法访问
*/

// 函数返回一个无命名变量
func sum(x int, y int) int {
	return x + y
}

// 参数类型一致的时候省略 变量类型 只要一个
func sub(x int, y int) int {
	return x - y
}

// 无参数列表和返回值
func printMyName() {
	fmt.Println("Bob Shen")
}

// 可变参数个数  类型相同
func show(args ...string) int {
	sum := 0
	for _, item := range args {
		fmt.Println(item)
		sum += 1
	}
	return sum
}

// 可变参数个数，类型不同
func diffTypeParam(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "type is int")
		case string:
			fmt.Println(arg, "type is string")
		case bool:
			fmt.Println(arg, "type is bool")
		default:
			fmt.Println(arg, "is an unknown type")
		}

	}
}

func main() {
	fmt.Printf("1+2= %v\n", sum(1, 2))
	fmt.Println("---------------------------------")
	fmt.Printf("2-1= %v\n", sub(2, 1))
	fmt.Println("---------------------------------")
	printMyName()
	fmt.Println("---------------------------------")
	show("c", "c++", "c#", "java", "python")
	fmt.Println("---------------------------------")
	diffTypeParam(1, "2", true, 21.5)
	fmt.Println("---------------------------------")
}
