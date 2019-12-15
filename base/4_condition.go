package main

import "fmt"

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
}
