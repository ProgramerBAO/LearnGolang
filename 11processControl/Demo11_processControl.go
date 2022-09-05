package main

import "fmt"

/*
	func condition() {

			条件
			if condition1{
				逻辑代码
			}else if condition2{
				逻辑代码
			}else if ...{
				逻辑代码 ...
			}else {
				逻辑代码 else
			}


	}
*/

func test01(n int) {
	score := n
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}

func advanceBranch() {
	/*
		if statement;condition{

		}
	*/
	if score := 88; score >= 60 {
		fmt.Println("成绩及格")
	}
}

func choose() {
	/*
		switch 表达式{
			case 表达式:
				业务逻辑
			case 表达式:
				业务逻辑
			case 表达式:
				业务逻辑
			default:
				默认业务逻辑
		}
	*/

	grade := "B"
	switch grade {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	default:
		fmt.Println("C")

	}
}

func main() {
	test01(90)
	fmt.Println("---------------------------------")
	advanceBranch()
	choose()
}
