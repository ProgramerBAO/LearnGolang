package main

import (
	"fmt"
	"sync"
	"time"
)

// 为了读写map

func SyncClassMap() {
	m := &sync.Map{}
	go func() {
		for {
			// 添加数据
			m.Store(1, 1)
			m.Store(2, 2)
			m.Store(3, 3)
			m.Range(func(key, value any) bool {
				fmt.Println(key, value)
				return true
			})
		}
	}()
	go func() {
		for {
			// 读取数据
			fmt.Println(m.Load(1))
			// 删除
			m.Delete(1)
		}
	}()
	time.Sleep(1000000)
}

func main() {
	SyncClassMap()
}
