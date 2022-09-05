package main

import "fmt"

func main() {

	var str1 = "str1"
	var str2 = "str2"
	var str3 = "str3"
	var str4 = "str4"

	var arr = [4]string{str1, str2, str3, str4}

	fmt.Println(arr[1])
	arr[1] = "111"
	fmt.Println(arr[1])
	/*
		这里其实看不出来指针的意义 
		指针的意义在于在方法调用的时候修改原始的值
		下面给出示例
	*/
	p1(str1)
	fmt.Println(str1)
	fmt.Println("------------分割线------------")
	p2(&str1)
	fmt.Println(str1)

	
}

func p1(p string){
	p = "这是不用指针的方法内做出的改变"
	fmt.Println(p)
}

func p2 (p *string){
	*p = "这是使用指针的方法内做出的改变"
	fmt.Println(*p)
}
