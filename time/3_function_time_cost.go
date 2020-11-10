package main

import (
	"fmt"
	"time"
)

// 统计函数花费时间，并打印
func main() {
	defer FuncTimeCost("DoSth")()
	DoSth()
}

func FuncTimeCost(logicName string) func() {
	start := time.Now()
	return func() {
		terminal := time.Since(start)
		fmt.Printf("func %s, cost time: %fs\n", logicName, terminal.Seconds())
	}
}

func DoSth() {
	fmt.Println("Start do sth.....")
	time.Sleep(2 * time.Second)
	fmt.Println("Do sth done.....")
}
