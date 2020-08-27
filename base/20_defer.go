package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

/**
defer: 延迟执行，用于进行资源释放等
一个方法中的多个defer，会“先进后出”的执行顺序，类似 栈的效果
*/
func main() {
	a := 1
	b := 2
	defer fmt.Println("front result:", add1(a, b)) // 3
	a = 3
	b = 4
	defer fmt.Println("last result:", add1(a, b)) // 7
	a = 5
	b = 6
	fmt.Println("final result:", add1(a, b)) // 11
}
