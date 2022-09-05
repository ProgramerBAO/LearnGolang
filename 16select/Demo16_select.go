package main

import "fmt"

func SelectTest01() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch1 <- "go"
	ch2 <- "c"
	ch3 <- "c++"

	select {
	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)
	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)
	case message3 := <-ch3:
		fmt.Println("ch3 received:", message3)
	default:
		fmt.Println("no data received")
	}

}

// 在多个发送接收通道进行选择
func main() {
	/*
		select{
		case expression1:
			code
		case expression2:
			code
		default:
			code
		}
	*/
	SelectTest01()

}
