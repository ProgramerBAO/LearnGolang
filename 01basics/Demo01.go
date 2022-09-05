package main //声明 main

import (
	"fmt" // 依赖包
	"os"
)

type User struct {
	UserId    string `json:"UserId"`
	SignH     string `json:"SignH"`
	PublicKey string `json:"PublicKey"`
}

func main() { // 声明 主函数
	fmt.Println("hello world")
	/*
		多行注释
	*/
	fmt.Println("nmsl")
	os.Getegid()
	bob := User{
		UserId:    "bob",
		SignH:     "11",
		PublicKey: "ss",
	}

	jack := User{
		UserId:    "bob",
		SignH:     "11",
		PublicKey: "ss",
	}

	fmt.Println([]User{bob, jack})
}
