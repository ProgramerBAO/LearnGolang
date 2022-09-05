package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 客户端
	// http包 创建一个客户端
	// client := http.Client{}
	client := new(http.Client)
	// get方法让客户端获取数据
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	res, _ := client.Do(req)
	// 将读取来的数据放在b里面
	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))
	// func (*http.Client).Get(url string) (resp *http.Response, err error)
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("client.Get err:", err)
		return
	}

	fmt.Println(resp.Header.Get("Content-Type"))
}
