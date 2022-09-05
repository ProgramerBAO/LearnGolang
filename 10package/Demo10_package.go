package main

import (
	"errors"
	// 这样就不需要使用 fmt. 来调用Println
	// . "fmt"
	"fmt"
	"strconv"
	// 这里导包 可以直接用 ./ 点斜杠  不知道为啥不会自动导入
	// p是别名
	//p ""
)

/*
	这玩意是默认存在的 这里显性地展示出来
	在包运行的时候 就会自动执行
	就和java的static一样（个人理解）
*/
func init() {
	fmt.Println("-------------init() -------------")
}

// package packageName
func main() {
	/*
		包的初始化顺序
		1 包级别的变量
		2 init()函数

		可以匿名导入
		import _ 包名

	*/
	// 直接Println 而不是fmt.Println  但是建议别用
	fmt.Println("If you shed tears when you miss the sun, you also miss the stars. ")
	fmt.Println("---------------------------------")
	// 这里使用了别名
	info, err := PersonInfo("Bob", 18)
	fmt.Println(info)
	fmt.Println(err)
}

// 首字母大写 可以被外部调用
func PersonInfo(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	} else if age <= 0 || age > 100 {
		return "", errors.New("wrong age")
	}
	// 这里强转了一下 age
	return "name:" + name + ",age:" + strconv.Itoa(age), nil
}
