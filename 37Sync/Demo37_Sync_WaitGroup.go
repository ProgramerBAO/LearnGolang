package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClassWaitGroup() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		time.Sleep(6 * time.Second)
		wg.Done()
		fmt.Println("6")
	}()

	go func() {
		time.Sleep(8 * time.Second)
		wg.Done()
		fmt.Println("8")
	}()
	wg.Wait()
	fmt.Println("*****************")
}

func main() {
	SyncClassWaitGroup()
}
