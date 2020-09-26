package main

import (
	"fmt"
	"time"
)

func main() {
	// 条件语句
	a := 2
	switch {
	case a == 1:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case a == 2:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case a == 3:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case a == 4:
		fmt.Println("4、case 条件语句为 true")
	case a == 5:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	nowTime := time.Now()
	switch nowTime.Weekday() {
	case time.Saturday:
		fmt.Println("Take a rest")
	case time.Sunday:
		fmt.Println("Take a rest")
	default:
		fmt.Println("Coding")
	}

	switch {
	case nowTime.Weekday() >= time.Monday && nowTime.Weekday() <= time.Friday:
		fmt.Println("weekday, you need work")
	default:
		fmt.Println("weekend~~")
	}

	// 带初始化的if (先初始化mm)
	if mm := "Mike"; mm == "Mike" {
		fmt.Println("Oh yeah, mm is Mike!")
	}
}
