package main

import "fmt"

func main() {
	//	range 关键字：用于for 循环中 迭代数组array、切片slice、通道channel、集合map 的元素
	// 数组和切片中，返回元素的索引和索引对应的值
	// 集合中，返回key-value对
	nums := []int{2, 3, 4}
	sum := 0
	// 当我们不想要这个index时，可以用"_"代替
	for _, num := range nums {
		sum += num
	}
	println("sum", sum)
	for i, num := range nums {
		println("遍历过程中：索引", i, "值", num)
	}

	// for map
	kvs := map[string]int{"A": 18, "B": 25}
	for k, v := range kvs {
		fmt.Printf("%s -> %d\n", k, v)
	}

	// for 枚举 Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		println(i, c)
	}
}
