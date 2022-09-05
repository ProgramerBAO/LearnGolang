package main

import (
	"fmt"
	"sync"
)

// bank
type Bank struct {
	balance int
}

type bank2 struct {
	balance int
	m       sync.Mutex // 互斥锁 还有读写锁(读多写少)
}

// 存钱
func (b *Bank) Deposit(amount int) {
	b.balance += amount
}

func (b *bank2) Deposit(amount int) {
	b.m.Lock()
	defer b.m.Unlock()
	b.balance += amount
}

// 查询余额
func (b *Bank) Balance() int {
	return b.balance
}

func (b *bank2) Balance() int {
	return b.balance
}

func test1(n int) int {
	var wg sync.WaitGroup
	b := &bank2{}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(1000)
			wg.Done()
		}()
	}
	wg.Wait()
	n, _ = fmt.Println(b.balance)
	return n
}

func main() {
	test1(10)
}
