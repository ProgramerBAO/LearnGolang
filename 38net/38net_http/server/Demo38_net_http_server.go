package main

import (
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	// http.ResponseWriter 是我们要返回的 req是前端传给我们的
	switch req.Method {
	case "GET":
		res.Write([]byte("我收到了给你返回GET"))
		break
	case "POST":
		// 读取前端传输过来的东西
		b, _ := io.ReadAll(req.Body)

		res.Write([]byte("我收到了给你返回POST"))
		res.Write(b)
		break
	}

}

func main() {
	// 注册路径，并且用handler处理数据
	http.HandleFunc("/test", handler)
	// 启动服务
	http.ListenAndServe(":9090", nil)
}
