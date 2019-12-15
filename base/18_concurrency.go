package main

import "time"

/**
Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式： go 函数名( 参数列表 )

Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
同一个程序中的所有 goroutine 共享同一个地址空间。
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	println("输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行")
}
