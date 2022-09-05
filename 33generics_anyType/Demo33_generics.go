package main

import "fmt"

// [T int | float64 | string] 用这种表达方式来表示对于可选的数据类型的约束
func same[T int | float64 | string](a, b T) bool {
	return a == b
}

// Person 泛型结构体
type Person[T any] struct {
	Name   string
	Gender T
}

// TMap 声明一个可变类型的map 泛型的map
type TMap[K string | int, V string | int] map[K]V

// TSlice 泛型的切片
type TSlice[S any] []S

// MyType02 ~的作用
type MyType02 interface {
	~int | ~int64
}

// MyInt 是 int的衍生类型
type MyInt int

func test04[T MyType02](t T) {
	fmt.Println(t)
}

func main() {
	b := same(1, 2)
	fmt.Println(b)

	p := Person[bool]{
		Name:   "bob",
		Gender: true,
	}
	fmt.Println(p)

	m := make(TMap[int, string])
	m[123] = "123"
	fmt.Println(m)

	s := make(TSlice[string], 6)
	s[5] = "456"

	// MyInt 是Int衍生而来 MyType 包含int且加上了~前缀 那么MyInt在这里就相当于Int 但是如果删除~ 那么就会报错
	test04[MyInt](123)
}
