package main

import (
	"fmt"
	"time"
)

// 最简单的 for死循环，每秒跑一次
// time.Sleep 背后也是创建timer
func main() {
	for {
		time.Sleep(time.Second)
		fmt.Println("run....")
	}
}
