package main

import (
	"fmt"
	"io"
	"net/http"
)

func handle(res http.ResponseWriter, req *http.Request) {
	// 这个就是对于请求的一些操作
	switch req.Method {
	case "GET":
		// 这句话是返回到前端的
		res.Write([]byte("我收到了，返回GET"))
	case "POST":
		// 读取前端过来的数据
		b, _ := io.ReadAll(req.Body)
		// 这句话是返回到前端的
		res.Write([]byte("我收到了，返回POST\n"))
		// 这里是请求头？？
		header := res.Header()
		header["test"] = []string{"test1, test2"}
		// 这是状态码
		res.WriteHeader(http.StatusBadRequest)
		res.Write(b)
		break

	}

}

func main() {
	/*
		注册路由
		xxxx/user ====> func1
		xxxx/help ====> func2
	*/
	//127.0.0.1:8080/user 就会触发这个
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		// writer : ====> 通过writer将数据返回给客户端
		// request: ====> 包含客户端发来的数据
		fmt.Println("这是user数据------")
		_, _ = io.WriteString(writer, "这是user数据------")
	})
	http.HandleFunc("/test", handle)

	fmt.Println("http server start")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http start failed,err:", err)
		return
	}
}
