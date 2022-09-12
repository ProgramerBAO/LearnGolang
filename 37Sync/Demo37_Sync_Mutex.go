package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClassMutex() {
	// 一定要用取地址符
	l := &sync.Mutex{}
	go LockFuncMutex(l)
	go LockFuncMutex(l)
	go LockFuncMutex(l)
	go LockFuncMutex(l)
	for {
	}
}

func LockFuncMutex(lock *sync.Mutex) {
	// 将它锁起来
	lock.Lock()
	fmt.Println("***********")
	// 休息一秒
	time.Sleep(1 * time.Second)
	// 解锁
	lock.Unlock()
}

func main() {
	SyncClassMutex()
}
