package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 0)
	go receive(ch2)

	// 主 goroutine 休眠 1s，保证调度成功
	time.Sleep(time.Second)
	for {
		// select 关键字：同时执行下面的每个case 的语句。
		select {
		case val := <-ch1: // 从 ch1 读取数据
			fmt.Printf("get value %d from ch1\n", val)
		case ch2 <- 2: // 使用 ch2 发送消息 （但是 receive 只会接收一次这个channel，后续消息的发送将被阻塞）
			fmt.Println("send value by ch2")
		case <-time.After(2 * time.Second): // 超时设置
			fmt.Println("Time out")
			return
		}
	}
}

func send(ch chan int, begin int) {
	for i := begin; i < begin+100; i++ {
		ch <- i
	}
}

func receive(ch <-chan int) {
	val := <-ch
	fmt.Println("receive:", val)
}
