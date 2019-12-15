package main

import "fmt"

// Go 语言切片是 对 数组 的抽象
// Go 数组 长度不可改变，在特定场景中，这样的不可变集合不适用
// 于是，Go提供了 一种灵活、功能强悍的内置类型切片（“动态数组”）
// 声明格式： var id []type
func main() {
	//var badArr1 [3]int = [3]int{1, 2, 3}
	badArr := [...]int{1, 2, 3}
	//  动态数组初始化方式一
	//var goodArr1 []int  = make([]int, 3)
	//  动态数组初始化方式二
	//goodArr2 := make([]int, 3)
	//  动态数组初始化方式三
	goodArr := [] int{1, 2, 3}

	println("普通数组")
	for i := 0; i < len(badArr); i++ {
		print(badArr[i], ",")
	}
	println()
	println("可变数组")
	for i := 0; i < len(goodArr); i++ {
		print(goodArr[i], ",")
	}
	println()
	goodArr = append(goodArr, 4, 5, 6)
	println("可变数组加了几个元素后")
	printSlice(goodArr)

	println()
	println()
	println("开始切片")
	s := [] int{1, 2, 3}
	println("这里有个切片s(len = cap = 3)：")
	printSlice(s)
	println("获取切片s后面两个元素：")
	printSlice(s[1:])
	println()
	println("获取切片s后面两个元素初始化为另一个切片s1：")
	s1 := s[1:]
	printSlice(s1)

	println()
	println("-------------------------------------------------")
	println("len() 和 cap()")

	var nums = make([]int, 3,5)
	printSlice(nums)

	println()
	println("-------------------------------------------------")
	println("空切片（nil）")

	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Printf("切片是空的")
	}
	println()
	println("-------------------------------------------------")
	println("append() 和 copy() 函数")
	append_and_copy()
}

func append_and_copy() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	//for i := 0; i < len(goodArr); i++ {
	//	print(goodArr[i], ",")
	//}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
