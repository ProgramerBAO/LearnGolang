package main

import "fmt"

func createSlice() {
	// 切片[]Type 数组[n]Type  Type 是数据类型的意思
	// 声明整形切片
	var numberList []int
	fmt.Println(numberList)

	// 声明一个空切片
	var numberListEmpty []int
	fmt.Println(numberListEmpty)

	// make声明方式 make([]Type,size,cap)
	numList := make([]int, 3, 6)
	fmt.Println(numList)
	// 指针 point: 是指向第一个切片元素对应底层数组元素的地址
	// 长度 size : 切片中的元素个数
	// 容量 cap  : 从切片的开始位置到底层数据的结尾位置

	arr := [...]string{"dddd", "yyds", "awsl", "yysy", "xswl"}
	var s1 = arr[1:4] // 数组变量[起始位置:结束位置] 左闭右开区间
	fmt.Println(arr)
	fmt.Println(s1)
	/*
		[]
		[]
		[0 0 0]
		[dddd yyds awsl yysy xswl]
		// 数组变量[起始位置:结束位置] 结束位置不包含，切片的数组个数应该是结束位置-起始位置
		var s1 = arr[1:4]
		[yyds awsl yysy]
	*/
}

func sliceLenAndCap() {
	s := make([]int, 3, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func outOfSlice() {
	s := make([]int, 3, 5)
	fmt.Println(s)
	// fmt.Println(s[10])
}

func emptySlice() {
	var numberList []int
	fmt.Println(numberList == nil)

	fmt.Println(len(numberList) == 0)
}

func modifySlice() {
	var arr = [...]string{"c", "c++", "c#"}
	s := arr[:]
	fmt.Println(arr)
	fmt.Println(s)

	s[0] = "go"
	fmt.Println(arr)
	fmt.Println(s)
}

func appendSliceData() {
	s := []string{"c", "c++", "c#"}
	// 追加一个元素
	s = append(s, "go")
	fmt.Println(s)
	fmt.Println(cap(s))
	// 追加两个元素
	s = append(s, "php", "html")
	fmt.Println(s)
	fmt.Println(cap(s))
	// 追加一个切片
	s = append(s, []string{"c", "c++", "c#"}...)
	fmt.Println(s)
	fmt.Println(cap(s))

	/*
		一旦发现追加的装不下，那么就会进行扩容，
		每次扩容是原有的2倍，扩容的方式是新建一片空间将原有的复制进去
	*/
}

func mSlice() {
	numList := [][]string{
		{"1", "c"},
		{"2", "c++"},
		{"3", "c#"},
	}
	fmt.Println(numList)
}

func main() {
	createSlice()
	fmt.Println("---------------------------------")
	sliceLenAndCap()
	fmt.Println("---------------------------------")
	outOfSlice()
	fmt.Println("---------------------------------")
	emptySlice()
	fmt.Println("---------------------------------")
	modifySlice()
	fmt.Println("---------------------------------")
	appendSliceData()
	fmt.Println("---------------------------------")
	mSlice()
}
