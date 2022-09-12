package main

import "fmt"

/*
	把一个接口类型指定为它的原始类型，
	接口，有的时候类型是不确定的，断言就是确定数据的类型，
	断言：下定论
*/
type User struct {
	Name   string
	Age    int
	Gender bool
}

type Student struct {
	User
}

func check(v interface{}) {
	// 这种行为叫做断言
	/*
		思考，在大型的项目开发过程中，由不同的程序员来编写不同的功能，
		就好像我要实现一个check接口，但是不知道怎么实现，就先预留这个check的接口，
		然后交由其他的程序员来开发这个接口，定义一些数据：例如 User
	*/
	// fmt.Println(v.(User).Name)

	// 在传输数据会发生变化，就像上文 student和User，差不多但不一样的时候，我们可以使用以下代码来进行操作
	// 这属于断言代码了
	switch v.(type) {
	case User:
		fmt.Println("User: " + v.(User).Name)
	case Student:
		fmt.Println("Student: " + v.(Student).Name)
	default:
		fmt.Println("oops, someting wrong!")
	}
}

func main() {
	u := User{
		Name:   "Bob",
		Age:    20,
		Gender: true,
	}
	check(u)

	s := Student{User{
		Name:   "Kenny",
		Age:    20,
		Gender: true,
	},
	}
	check(s)
}
