package main

import (
	"fmt"
	"sync"
	"time"
)

//通知和锁

func SyncClassCond() {
	// 首先传个锁给它
	c := sync.NewCond(&sync.Mutex{})

	go func() {
		c.L.Lock()
		fmt.Println("lock1")
		// 这里可以用flag来进行控制 特别是使用Signal()进行解锁的时候
		c.Wait()
		fmt.Println("unlock1")
		c.L.Unlock()
	}()
	go func() {
		c.L.Lock()
		fmt.Println("lock2")
		c.Wait()
		fmt.Println("unlock2")
		c.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	//// 上锁
	//c.L.Lock()
	//// 可以执行等待
	//c.Wait()
	//// 解锁
	//c.L.Unlock()
	//// 通知完成一个等待
	//c.Signal()
	// 广播所有的锁解锁
	c.Broadcast()
	time.Sleep(2 * time.Second)
}

func main() {
	SyncClassCond()
}
