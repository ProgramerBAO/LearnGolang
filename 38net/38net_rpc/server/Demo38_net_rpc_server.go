package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Num int
}

func (s *Server) Add(req Req, res *Res) error {
	res.Num = req.NumOne + req.NumTwo
	return nil
}
func main() {
	// 注册服务
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
	}
	// 开启服务
	http.Serve(l, nil)
}
