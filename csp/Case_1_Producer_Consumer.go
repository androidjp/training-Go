package main

import (
	"fmt"
	"time"
)

// chan<- 表示channel发送端
// <-chan 表示channel接收端
// 生产者 和 消费者，通过 go 语法，都变成 goroutine
// main函数 由 主goroutine 启动
// 当 主goroutine 即 main 函数 执行结束，整个 Go程序也会直接执行结束，无论是否存在其他未执行完的 goroutine

func Producer(begin, end int, queue chan<- int) {
	for i := begin; i < end; i++ {
		fmt.Println("produce:", i)
		queue <- i
	}
}

func Consumer(queue <-chan int) {
	for val := range queue {
		fmt.Println("consume:", val)
	}
}

func main() {
	// 制作一条 channel
	queue := make(chan int)
	defer close(queue)
	// 三个生产者
	for i := 0; i < 3; i++ {
		go Producer(i*5, (i+1)*5, queue)
	}
	// 一个消费者
	go Consumer(queue)
	// 睡 1s, 避免 主goroutine 结束程序
	time.Sleep(time.Second)
}
