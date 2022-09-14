package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Num int
}

func main() {
	req := Req{
		NumOne: 2,
		NumTwo: 3,
	}
	var res Res
	// 这里就好像与服务其建立起了通信，传递书库
	client, err := rpc.DialHTTP("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err)
	}
	// 远程调用
	client.Call("Server.Add", req, &res)
	fmt.Println(res)

}
