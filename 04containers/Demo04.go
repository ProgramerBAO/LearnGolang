package main

import "fmt"

// 数组
func test01() {
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)

}

func test02() {
	// 声明时对数组进行初始化
	var arr1 = [5]int{6, 7, 8, 9, 10}
	fmt.Println(arr1)

	// 短声明
	arr2 := [5]int{23, 45, 67, 89, 99}
	fmt.Println(arr2)

	// 部分初始化
	arr3 := [5]int{89, 99}
	fmt.Println(arr3)

	// 通过指定索引对某几个数组进行初始化
	arr4 := [5]int{1: 22, 4: 33}
	fmt.Println(arr4)

	// 让系统自动识别元素个数
	arr5 := [...]int{23, 45, 67, 89, 99, 88, 77}
	fmt.Println(arr5)
}

func test03() {
	// 数组的长度是数组长度的一部分 [3]int and [5] is different
	arr1 := [3]int{25, 20, 23}
	arr2 := [5]int{23, 45, 67, 89, 99}
	fmt.Printf("type of arr1 %T\n", arr1)
	fmt.Printf("type of arr2 %T\n", arr2)
	/*
		type of arr1 [3]int
		type of arr2 [5]int
	*/
}

func test04() {
	// 创建多维数组
	arr := [...][2]string{
		{"1", "dddd"},
		{"2", "yyds"},
		{"3", "awsl"},
	}
	fmt.Println(arr)
}

func arrLength() {
	// len() 返回数组元素个数
	arr := [...]string{"dddd", "yyds", "awsl"}
	fmt.Println("数组的长度是:", len(arr))
}

func showArr() {
	// for range 用于遍历数组
	arr := [...]string{"dddd", "yyds", "awsl"}
	for index, value := range arr {
		fmt.Printf("arr[%d]=%s\n", index, value)
	}

	// 可以使用下划线代替不使用的变量
	for _, value := range arr {
		fmt.Printf("value=%s\n", value)
	}
}

func arrByValue() {
	arr := [...]string{"dddd", "yyds", "awsl"}
	copy := arr
	copy[0] = "golange"
	fmt.Println(arr)
	fmt.Println(copy)
	/*
		[dddd yyds awsl]
		[golange yyds awsl]
		这个结果可以说明。copy开辟了一个新的空间，而不是想python一样是用”指针“指向空间的。
		是值类型而不是引用类型
	*/
}

func main() {
	test01()
	test02()
	test03()
	test04()
	arrLength()
	showArr()
	arrByValue()
}
