package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	// 循环监听
	for {
		TCPConn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(TCPConn)
	}

}

func handleConnection(conn *net.TCPConn) {
	for {
		// 创建一个缓存去
		buf := make([]byte, 1024)
		// 进行读取 吧东西放入缓存区 先不管怎么实现的 以下操作之后 buf中就有内容了
		n, err := conn.Read(buf)
		if err != nil {
			// 来判断有没有内容
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + string(buf[0:n]))
		str := "收到了" + string(buf[0:n])
		// 这是往回写的操作
		conn.Write([]byte(str))
	}

}
