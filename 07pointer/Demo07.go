package main

import "fmt"

func pointer() {
	/*
		指针不能指向一个具体的值
			var a int
			a = 123
			var b *int
			b = a  这样是不行的
			b = &a 这个时候 b指向了a的地址
			*b = 999  这个时候 a = 999
	*/
	// 通过自定义普通变量获取指针
	x := "go"
	// 获取地址
	ptr := &x
	fmt.Println("x=", x)
	fmt.Println("ptr=", ptr)
	fmt.Println("&x=", &x)
	fmt.Println("*ptr=", *ptr)
	fmt.Println("---------------------------------")

	// new是先创建好指针并分配内存，再写入指针的值
	ptr2 := new(string)
	*ptr2 = "c"
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	// 先声明一个指针变量，再从其他变量获取内存地址指针变量
	x2 := 1
	var p *int = &x2 // *int 因为 x2 是int类型的
	fmt.Println(p)
	fmt.Println(*p)

	// & 可以通过变量获取该变量的地址
	// * 赋在值的左边(*p) 指该指针指向的变量(p是变量地址)
}

func pointerType() {
	myStr := "go"
	myInt := 1
	myBool := false
	myFloat := 3.2
	fmt.Printf("type of &myStr is:%T\n", &myStr)
	fmt.Printf("type of &myInt is:%T\n", &myInt)
	fmt.Printf("type of &myBool is:%T\n", &myBool)
	fmt.Printf("type of &myFloat is:%T\n", &myFloat)
}

func zeroPointer() {
	x := "go"
	var ptr *string
	fmt.Println("ptr is ", ptr)
	// 声明但是没有初始化
	ptr = &x
	fmt.Println("ptr is ", ptr)

}

func changeByPointer(value *int) {
	// value是地址 *value是值
	*value = 200
}

func test01() {
	x3 := 99
	p3 := &x3
	fmt.Println("执行changeByPointer函数之前*p3是", *p3)
	fmt.Println("执行changeByPointer函数之前p3是", p3)
	changeByPointer(p3)
	fmt.Println("执行changeByPointer函数之后*p3是", *p3)
	fmt.Println("执行changeByPointer函数之后p3是", p3)
}

func changeSlice(value []int) {
	// 这里是传递的值进来
	value[0] = 200
}

func changeArray(value *[3]int) {
	// 这里是传递的地址
	(*value)[1] = 200
}

func test02() {
	x4 := [3]int{10, 20, 30}
	changeSlice(x4[:])
	fmt.Println(x4)
	changeArray(&x4)
	fmt.Println(x4)
}

func main() {
	// go语言中不支持指针运算
	pointer()
	fmt.Println("---------------------------------")
	pointerType()
	fmt.Println("---------------------------------")
	zeroPointer()
	fmt.Println("---------------------------------")
	test01()
	fmt.Println("---------------------------------")
	test02()
}
