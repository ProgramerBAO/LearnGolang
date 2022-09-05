package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "1001")

	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	fmt.Println("client connected service")

	sendData := "hello"
	// 向服务器发送数据
	cnt, err := conn.Write([]byte(sendData))

	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	fmt.Println("client=====>server,长度", cnt, "，数据:", sendData)

	// 接收服务器返回数据
	buf := make([]byte, 1024)

	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("client<=====server,长度", cnt, "，数据:", string(buf[0:cnt]))

	conn.Close()
}
