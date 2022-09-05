package main

import (
	"fmt"
	"reflect"
)

// 类似于java中的类 这个更加清晰
type person struct {
	name  string
	age   int
	hight int
}

type person2 struct {
	name       string
	age, hight int
}

// 匿名字段 只给类型不命名字段 每个类型只能有一个
type person3 struct {
	string
	int
}

// 嵌套结构体
type inter struct {
	name string
	age  int
}

type outer struct {
	id    int
	inter inter
}

func test01() {
	Bob := person{
		name:  "bob",
		age:   18,
		hight: 190,
	}
	fmt.Println(Bob)
}

func test02() {
	// 按照字段顺序进行初始化
	Alice := person2{
		"alice", 18, 189,
	}
	fmt.Println(Alice)
}

func test03() {
	// 匿名结构体
	Jack := struct {
		name  string
		age   int
		hight int
	}{
		name:  "Jack",
		age:   20,
		hight: 178,
	}
	fmt.Println(Jack)
}

func test04() {
	// 不赋值的时候的默认值
	var nobody = person{}
	fmt.Println(nobody)
}

func test05() {
	// 声明指定字段
	Candy := person{
		name: "Candy",
	}
	fmt.Println(Candy)
}

func test06() {
	// 对于结构体中字段的单独访问与赋值
	Douglas := person{
		name:  "Douglas",
		age:   18,
		hight: 24,
	}
	fmt.Println("name:", Douglas.name)
	fmt.Println("age:", Douglas.age)
	fmt.Println("hight", Douglas.hight)
	Douglas.age = 19
	fmt.Println("age:", Douglas.age)
}

func test07() {
	Douglas := &person{
		name:  "Douglas",
		age:   18,
		hight: 24,
	}
	fmt.Println("name:", (*Douglas).name)
	fmt.Println("age:", (*Douglas).age)
	fmt.Println("hight", (*Douglas).hight)
	fmt.Println("age:", Douglas.age) // 解引用的方式访问
}

func test08() {
	Kenny := person3{
		"kenny",
		12,
	}
	fmt.Println("Kenny:", Kenny)
	fmt.Println("Kenny string", Kenny.string)
	fmt.Println("Kenny int", Kenny.int)
}

func test09() {
	Penny := outer{
		id: 12345,
		inter: inter{
			name: "Penny",
			age:  18,
		},
	}

	fmt.Println("Penny:", Penny)
	// 字段的提升
}

func test10() {
	// 比较
	Kenny01 := person3{
		"kenny",
		12,
	}

	Kenny02 := person3{
		"kenny",
		12,
	}

	fmt.Println(Kenny01 == Kenny02)
	fmt.Println(reflect.DeepEqual(Kenny01, Kenny02))

}

// 结构体可以有方法
func (l person) showPersonInfo(n int) {
	fmt.Println("name:", l.name)
	fmt.Println("age", l.age)
	fmt.Println("age+n", l.age+n)
}

// 我不明白这和上面有啥区别
func (p *person) addAge(n int) {
	p.age = p.age + n
	fmt.Println("addAge:", p.age)
}

func test11() {
	Bob := person{
		name:  "bob",
		age:   18,
		hight: 190,
	}
	fmt.Println("showPersonInfo(8)----------------")
	Bob.showPersonInfo(8)
	fmt.Println("addAge(10)-----------------------")
	Bob.addAge(10)
}

func main() {
	test01()
	fmt.Println("---------------------------------")
	test02()
	fmt.Println("---------------------------------")
	test03()
	fmt.Println("---------------------------------")
	test04()
	fmt.Println("---------------------------------")
	test05()
	fmt.Println("---------------------------------")
	test06()
	fmt.Println("---------------------------------")
	test07()
	fmt.Println("---------------------------------")
	test08()
	fmt.Println("---------------------------------")
	test09()
	fmt.Println("---------------------------------")
	test10()
	fmt.Println("---------------------------------")
	test11()
	fmt.Println("---------------------------------")
}
