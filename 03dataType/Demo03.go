package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Integer 有符号整形
func Integer() {
	var num8 int8 = math.MaxInt8
	var num16 int16 = math.MaxInt16
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8的类型是 %T,num8的大小是 %d,num8的取值是  %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T,num16的大小是 %d,num16的取值是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T,num32的大小是 %d,num32的取值是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T,num64的大小是 %d,num64的取值是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T,num的大小是 %d,num的取值是 %d\n", num, unsafe.Sizeof(num), num)
}

// 无符号整型
func unsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint
	fmt.Printf("num8的类型是 %T, num8的大小 %d, num8是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T, num16的大小 %d, num16是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T, num32的大小 %d, num32是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T, num64的大小 %d, num64是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T, num的大小 %d, num是 %d\n", num, unsafe.Sizeof(num), num)
}

func showFloat() {
	var num1 float32 = math.MaxFloat32 //小数点后六位
	var num2 float64 = math.MaxFloat64 //小数点后15位
	fmt.Printf("num1的类型是 %T, num1是%g\n", num1, num1)
	fmt.Printf("num2的类型是 %T, num2是%g\n", num2, num2)
}

func main() {
	Integer()
	println("---------------------------------------")
	unsignedInteger()
	println("---------------------------------------")
	showFloat()
}
