package main

import (
	"fmt"
	"time"
)

// 如果是默认不给时区，那么，就会导致，Parse得到的一定是0时区时的毫秒数。
func main() {
	layout := "2006-01-02 15:04:05 -0700" //  -0700

	now := time.Now()
	fmt.Println("当前时间")
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.Format(layout))

	fmt.Println("-------------------------------------")

	timeStr := "2020-11-03 10:43:01 +0800"
	fmt.Println("接收到的时间")
	fmt.Println(timeStr)
	t1, _ := time.Parse(layout, timeStr)
	fmt.Println(t1.Unix())
	fmt.Println("当前时间是否在给定的时间之前？", now.Before(t1))
}
