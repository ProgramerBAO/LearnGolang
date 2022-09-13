package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":9999")
	// 建立连接信道 Dial 拨号上网
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	reader := bufio.NewReader(os.Stdin)
	// 长链接
	for true {
		by, _, _ := reader.ReadLine()
		conn.Write(by)
		rb := make([]byte, 1024)
		// 将返回的信息写入
		n, _ := conn.Read(rb)
		fmt.Println(string(rb[0:n]))
	}
}
