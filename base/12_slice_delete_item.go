package main

// Go 语言切片是 对 数组 的抽象
// Go 数组 长度不可改变，在特定场景中，这样的不可变集合不适用
// 于是，Go提供了 一种灵活、功能强悍的内置类型切片（“动态数组”）
// 声明格式： var id []type
// 往切片中删除某个元素
func main() {
	//var badArr1 [3]int = [3]int{1, 2, 3}
	badArr := [...]int{1, 2, 3}
	//  动态数组初始化方式一
	//var goodArr1 []int  = make([]int, 3)
	//  动态数组初始化方式二
	//goodArr2 := make([]int, 3)
	//  动态数组初始化方式三
	goodArr := []int{1, 2, 3}

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
	println("此时的动态数组：")
	printSlice(goodArr)
	goodArr = deleteItemForSlice(goodArr, 5)
	println("此时的动态数组：")
	printSlice(goodArr)
}

/*
删除某个元素
*/
func deleteItemForSlice(list []int, x int) []int {
	for i, item := range list {
		if item == x {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
