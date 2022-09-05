package main

import "fmt"

func PrintChan(ch chan string) {
	ch <- "go"
}

/*
	waitGroup 等待一组任务结束在执行其他业务逻辑
	Add() 初始值是0，累加子协程的数量
	Done() 当某个子协程完成后，计数器减一
	Wait() 阻塞当前协程，直到实例中的计数器归零
*/

func main() {
	/*
		通道声明
		var channelName chan channelType
	*/
	// var ch chan string
	// ch = make(chan string)
	// ch:=make(chan string)
	/*
		发送数据
		channelName <- date

		接收数据
		value := <- channelName

		关闭通道
		close(channelName)

		可以事先判断一下通道状态
		value, ok := <-channelName

		if ok == true {
			close(channelName)
		}

		通道可以有容量，并且可以设置只读只写 稍后再学
		利用这种机制做锁
	*/
	ch := make(chan string)

	fmt.Println("学习课程:")
	// 开启协程
	go PrintChan(ch)
	// 从通道中接收数据
	rec, ok := <-ch
	// 打印从通道中接受的数据
	fmt.Println(rec)
	fmt.Println(ok)
	if ok {
		close(ch)
	}

}
