package main

import "fmt"

func main() {

	//	声明数组
	var balance [10] float32
	balance[0] = 12
	println("balance数组长度", len(balance))
	println("balance数组cap", cap(balance))
	println("balance数组指针", &balance)

	var actors [10] string
	actors[0] = "XiaoMing"
	println("actors数组长度", len(actors))
	println("actors数组cap", cap(actors))
	println("actors数组指针", &actors)

	//	初始化数组(括号里的参数个数不得大于 给定的数组长度)
	// 报错：因为前面已经声明 balance 为长度是 10 的float32 数组了
	//balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var arr1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	println(len(arr1))
	// 赋值的数组形状一样，所以不会报错
	// 这里可以不指定数组长度，默认按照括号中的元素个数来设置数组大小
	arr1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	println(len(arr1))

	println()
	println()
	println("遍历数组（保留2位小数）：")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%.2f,", arr1[i])
	}
}
