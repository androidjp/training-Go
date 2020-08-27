package main

import (
	"fmt"
	"time"
)

func main() {
	worker1Callback := make(chan string)
	ch2 := make(chan string)
	defer close(ch2)
	defer close(worker1Callback)

	// Admin 让  worker1 干活
	go worker1DoSth(worker1Callback, 0)
	go worker1DoSth(worker1Callback, 0)
	go worker1DoSth(worker1Callback, 1)
	go worker1DoSth(worker1Callback, 2)
	//go tellAdminMonther(ch2)


	for {
		select {
		case response := <-worker1Callback:
			go tellAdminMonther(ch2)
			ch2 <- response
		case <-time.After(1 * time.Second):
			go tellAdminMonther(ch2)
			ch2 <- "time out! Mom"
			return
		}
	}
}

func tellAdminMonther(ch chan string) {
	msg := <-ch
	fmt.Println("Mother receive the msg: ", msg)
}

func worker1DoSth(ch chan string, command int) {
	switch command {
	case 0:
		fmt.Println("worker1: coding")
		ch <- "I'm coding now"
	case 1:
		fmt.Println("worker1: buying coffee")
		ch <- "I'm buying coffee"
	default:
		fmt.Println("worker1: clean the office")
		ch <- "I'm cleaning the office"
	}
}
