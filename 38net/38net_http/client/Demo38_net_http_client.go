package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 创建客户端
	client := new(http.Client)
	// 构建请求
	req, _ := http.NewRequest("GET", "http://localhost:9090/test", nil)
	res, _ := client.Do(req)
	// 查看结果
	body := res.Body
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
	//
}
