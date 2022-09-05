package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建监听
	fmt.Println("create a listener")
	ip := "127.0.0.1"
	// 这里不用冒号包裹
	port := 9001
	address := fmt.Sprintf("%s:%d", ip, port)
	/*
		func Listen(network, address string) (Listener, error)
	*/
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("oops", err)
		return
	}
	fmt.Println("监听ing")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("connected")
	// 创建一个容器，用于接收数据
	// 创建一个切片
	buf := make([]byte, 1024)
	// cnt：这正读取client发来的数据的长度
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("client=====>server,长度", cnt, "，数据:", string(buf))

	// 将数据转成大写 服务器对客户端请求进行响应
	upperData := strings.ToUpper(string(buf))
	// 写入连接 这里不需要：冒号
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("client<=====server,长度", cnt, "，数据:", upperData)

	// 关闭连接
	conn.Close()
}
