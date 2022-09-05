package main

import (
	"fmt"
)

// map是引用类型
func createMap() {
	// make(map[keyType]ValueType)
	// map是无序的
	// 创建相同数据类型的键值对
	steps := make(map[string]string)
	fmt.Println(steps)

	// 键值对类型不同
	scores := make(map[string]int)
	fmt.Println(scores)

	// 通过字面值创建map
	var steps2 = map[string]string{
		"1": "c",
		"2": "go",
		"3": "python",
	}
	fmt.Println(steps2)

	fmt.Println()

	steps3 := map[string]string{
		"1": "c",
		"2": "go",
		"3": "python",
	}
	fmt.Println(steps3)

}

func mapAddData() {
	m := make(map[string]string)
	// 添加元素
	m["1"] = "c"
	fmt.Println(m)
	// 更新元素
	m["1"] = "go"
	fmt.Println(m)
	// 获取元素
	fmt.Println(m["1"])
	// 删除元素
	delete(m, "1")
	fmt.Println(m)

	// 判断键值是否存在 value,ok := map[key]
	steps3 := map[string]string{
		"1": "c",
		"2": "go",
		"3": "python",
	}
	fmt.Println(steps3)
	v3, ok := steps3["1"]
	fmt.Println(v3) // 如果有值 那么就返回一个值
	fmt.Println(ok) // 如果有值 在返回值的同时给ok赋值为true

	v4, ok := steps3["4"]
	fmt.Println(v4) // 如果没有值 那么就没有返回值
	fmt.Println(ok) // 如果没有值 同时给ok赋值为true

	// 遍历map for range
	for key, value := range steps3 {
		fmt.Printf("key:%s,value:%s\n", key, value)
	}
	// len()获取map长度
	fmt.Println(len(steps3))

}

func mapByReference() {
	// map是引用类型

	steps3 := map[string]string{
		"1": "c",
		"2": "go",
		"3": "python",
	}
	fmt.Println(steps3)
	newStep3 := steps3
	fmt.Println("---------------------------------")
	newStep3["1"] = "java"
	newStep3["2"] = "html"
	newStep3["3"] = "php"
	fmt.Println("---------------ssss------------------")
	fmt.Println(steps3["4"] == "4")
	fmt.Println("----------------ssss-----------------")
	fmt.Println(newStep3)

}

func main() {
	createMap()
	fmt.Println("---------------------------------")
	mapAddData()
	fmt.Println("---------------------------------")
	mapByReference()
	fmt.Println("---------------------------------")
	count := 5
	for count != 0 {
		fmt.Println(count)
		count = count - 1
	}

}
