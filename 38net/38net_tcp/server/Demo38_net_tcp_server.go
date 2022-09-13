package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建tcp地址 协议为tcp  地址为localhost:9999  返回的是一个tcpAddr的指针
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":9999")
	// 监听
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	// 返回一个tcp的连接 for循环实时监听
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
	for true {
		// 把信息读取出来
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + string(buf[0:n]))
		str := "收到了：" + string(buf[0:n])
		// 写回去，其实conn就像是一个管道，链接了服务端和客户端
		conn.Write([]byte(str))
	}
}
