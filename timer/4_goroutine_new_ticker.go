package main

import (
	"fmt"
	"time"
)

// 通过goroutine 加上 recover，来解决主动中止调度出现的panic问题
func main() {
	go scheduleJobA()
	time.Sleep(2 * time.Second)
}

func scheduleJobA() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
		}
	}()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("run...")
			ticker.Reset(2 * time.Second)
			ticker.Stop()
		}
	}
}
