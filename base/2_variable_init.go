package main

import "fmt"

func main() {
	// 声明变量一般用 var 关键字
	// 格式：var identifier type
	// 可以一次声明多个变量
	// 格式：var id1, id2 type
	var a string = "Mike"
	fmt.Println(a) // Mike

	var b, c int = 1, 2
	fmt.Print(b, c) // 1 2

	fmt.Println("变量声明方式一：指定变量类型，如果未初始化，则变量默认零值")
	// 变量声明方式一：指定变量类型，如果未初始化，则变量默认零值
	// var v_name v_type
	// v_name = value
	// 声明一个变量并初始化
	var a1 = "Jay"
	fmt.Println(a1) // Jay

	// 没有初始化就为零值
	var b1 int
	fmt.Println(b1) // 0

	// bool 零值为 false
	var c1 bool
	fmt.Println(c1) // false


	var intVal int

	//intVal :=1 // 这时候会产生编译错误

	intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句

	fmt.Println(intVal, intVal1)

	var i = 7
	var j = i
	fmt.Println(&i, &j)
}
