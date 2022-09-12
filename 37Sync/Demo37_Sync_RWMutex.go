package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClassRWMutex() {
	// 一定要用取地址符
	l := &sync.RWMutex{}
	go LockFuncWMutex(l)
	go LockFuncWMutex(l)
	go LockFuncWMutex(l)
	go LockFuncWMutex(l)
	go LockFuncRMutex(l)
	go LockFuncRMutex(l)
	go LockFuncRMutex(l)
	go LockFuncRMutex(l)
	go LockFuncRMutex(l)
	for {
	}
}

func LockFuncWMutex(lock *sync.RWMutex) {
	// 将它锁起来,写的时候会排斥其他的写锁，就是同时智能一个写
	lock.Lock()
	fmt.Println("***********")
	// 休息一秒
	time.Sleep(1 * time.Second)
	// 解锁
	lock.Unlock()
}

func LockFuncRMutex(lock *sync.RWMutex) {
	// 将它锁起来,在读取的时候不会排斥其他的读取锁，但是会排斥其他的写锁，能读的时候不能写，但能读，还能好多人一起读
	lock.RLock()
	fmt.Println("--------------")
	// 休息一秒
	time.Sleep(1 * time.Second)
	// 解锁
	lock.RUnlock()
}

func main() {
	SyncClassRWMutex()
}
