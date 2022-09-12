package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发池

func SyncClassPool() {
	p := &sync.Pool{}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)
	p.Put(5)
	p.Put(6)
	p.Put(7)

	for i := 0; i < 8; i++ {
		// 一定是先把有的值取完再取空值
		fmt.Println(p.Get())
		time.Sleep(1 * time.Second)
	}
}

func main() {
	SyncClassPool()
}
