package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//---------------------------------------
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area) // 面积为：50
	println()
	println(a, b, c) // 1 false str

	//-----------------------------------------
	const (
		a_1         = "abc"
		a_len       = len(a_1)
		a_type_size = unsafe.Sizeof(a_1)
	)
	println(a_1, a_len, a_type_size)
}
