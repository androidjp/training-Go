package main

import (
	"fmt"
	"time"
)

func main() {
	mChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(mChan)
	})

	select {
	case _, ok := <-mChan:
		if !ok {
			fmt.Println("已经关了")
			break
		}
		fmt.Println("哈哈哈")
	}
}
