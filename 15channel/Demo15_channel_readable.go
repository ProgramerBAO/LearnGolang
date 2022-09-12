package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	// 只可读
	var readc <-chan int = c1
	// 只可写
	var writec chan<- int = c1

	writec <- 1
	fmt.Println(<-readc)
	// 关闭channel
	close(c1)
}
