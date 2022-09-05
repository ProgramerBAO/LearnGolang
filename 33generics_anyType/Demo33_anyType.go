package main

import (
	"fmt"
	"strconv"
)

// A or T 不过只是个名字罢了 后面any是一种约束
type My[A any] struct {
	Name A
}

// generics 泛型 我更喜欢叫它 any type
func test01(i int) string {
	// 把int类型的i 强转成string
	return strconv.Itoa(i)
}
func test02[T any](i T) T {
	// 把int类型的i 强转成string
	return i
}

type MyType interface {
	getValue() string
}

func test03[T MyType](t T) {
	fmt.Println(t.getValue())
}

type my struct {
	Name string
}

func (m my) getValue() string {
	return m.Name
}

func main() {
	fmt.Println(test01(123))
	fmt.Println(test02[string]("1234") + "46")
	fmt.Println(test02[int](1234) + 56)

	m := My[string]{
		Name: "Bob",
	}

	fmt.Println(m)

}
