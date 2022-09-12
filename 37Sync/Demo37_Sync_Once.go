package main

import (
	"fmt"
	"sync"
)

func SyncClassOnce() {
	// 只执行一次
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println(i)
		})
	}
	for {
	}
}

func main() {
	SyncClassOnce()
}
