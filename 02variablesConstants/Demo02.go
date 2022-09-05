package main

import "fmt"

const (
	BEIJING   = 0
	SHANGHAI  = 1
	SHENZHENG = 2
)

const ( // 从0开始依次递增 可以设置步长 这里是10
	HANGZHOU = 10 * iota
	CHONGQING
	CHENGDU
)

func main() {
	// var 关键字
	// 声明变量，默认值是0
	var a int
	fmt.Printf("a=%d,a的类型是:%T\n", a, a)

	// 声明变量并赋值
	var b int = 100
	fmt.Printf("b=%d,c的类型是:%T\n", b, b)

	// 初始化的时候去掉数据类型
	var c = 100

	fmt.Printf("c=%d,c的类型是:%T\n", c, c)

	var st = "ccc"
	fmt.Printf("st=%s,st的类型是:%T\n", st, st)

	// 短声明
	e := 100
	fmt.Printf("e=%d,e的类型是:%T\n", e, e)
	s := "hello"
	fmt.Printf("s=%s,s的类型是:%T\n", s, s)

	// 可以同时声明多个变量 可以是同类型也可以是不同的
	var xx, yy = 100, 200
	var aa, bb = 300, "hello"
	fmt.Printf("xx=%d,xx的类型是:%T\n", xx, xx)
	fmt.Printf("yy=%d,yy的类型是:%T\n", yy, yy)
	fmt.Printf("aa=%d,aa的类型是:%T\n", aa, aa)
	fmt.Printf("bb=%s,bb的类型是:%T\n", bb, bb)

	var (
		nn = 100
		mm = "hello"
	)
	fmt.Printf("nn=%d,nn的类型是:%T\n", nn, nn)
	fmt.Printf("mm=%s,mm的类型是:%T\n", mm, mm)

	// 常量
	const constant = 10

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHENG=", SHENZHENG)

	// iota 常量计数器 枚举时使用
	fmt.Println("HANGZHOU=", HANGZHOU)
	fmt.Println("CHONGQING=", CHONGQING)
	fmt.Println("CHENGDU=", CHENGDU)
}
