package main

// 之前一直觉得方法和函数差不多
import "fmt"

/*
	// 这是方法
	func(t Type) methodName(parameterList) resultList{

	}

	// 这是函数
	func functionName(parameterList)(resultList){
		// 函数体
		参数列表和返回值列表不是必须的
	}

	在go中相同名字的方法可以定义在不同的类型上，而想同类型的函数是不行的

	指针接收器 可以改变数据
*/

type Person struct {
	name string
	age  int
}

func (person Person) PrintInfo() {
	fmt.Println("name:", person.name)
	fmt.Println("age:", person.age)
}

func main() {
	person := Person{
		name: "Bob",
		age:  18,
	}
	person.PrintInfo()
	fmt.Println("---------------------------------")
}
