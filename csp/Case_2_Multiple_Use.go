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

	/**
	select 语句分别从 3 个 case 中选取返回的 case 进行处理，当有多个 case 语句同时返回时，select 将会随机选择一个 case 进行处理。
	如果 select 语句的最后包含 default 语句，该 select 语句将会变为非阻塞型，即当其他所有的 case 语句都被阻塞无法返回时，select 语句将直接执行 default 语句返回结果。
	在上述例子中，我们在最后的 case 语句使用了 <-time.After(2 * time.Second) 的方式指定了定时返回的 channel，这是一种有效从阻塞的 channel 中超时返回的小技巧。
	 */
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
