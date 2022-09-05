package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	// dial是拨的意思
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("出错啦！！")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := reader.ReadLine()
		conn.Write(bytes)
		// 读取服务器端返回的信息
		rb := make([]byte, 1024)
		rn, _ := conn.Read(rb)
		fmt.Println(string(rb[0:rn]))
	}

}
