package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func file() {
	// 文件路径  如何操作  权限
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	// 拷贝
	f2, _ := os.OpenFile("./testCopy.txt", os.O_CREATE|os.O_RDWR, 0777)
	io.Copy(f2, f)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// 这是一个装载数据的容器，并且注意一个汉字占了多少位
		b := make([]byte, 1024)
		// n 表示容量 当读到最后读不出的时候会返回err
		n, err := f.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b), n)
	}
	// 写入操作
	// 在这里要用双引号
	f.WriteString("\n")
	f.Write([]byte("我读完了"))
	fmt.Println("-------------------我是分割线-------------------")
	//通过bufio来读取
	f.Seek(0, io.SeekStart)
	bufReader := bufio.NewReader(f)
	for {
		// \n要用单引号包裹
		str, err := bufReader.ReadString('\n')
		fmt.Println(str)
		if err != nil {
			fmt.Println(err)
			break
		}

	}
	//通过bufio来写
	f.Seek(0, io.SeekStart)
	w := bufio.NewWriter(f)
	r := bufio.NewReader(f)
	n := 0
	for {
		n++
		str, err := r.ReadString('\n')
		w.WriteString(strconv.Itoa(n) + " " + str)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	f.Seek(0, 0)
	w.Flush()

}

func main() {
	file()
}
