package main

import "fmt"

func main() {
	println("printInt:", printInt())
	println("add(1,2):", add(1, 2))
	s := "hello world"
	println("s的长度:", len(s))

	baseSalary := 8000.0
	var rate float32 = 1.2
	var res float64 = calculateFinalSalary(baseSalary, rate)
	fmt.Printf("原来工资为%f,晋升后的工资为：%f\n", baseSalary, res)

	x1 := "Mike"
	x2 := "David"
	res1, res2 := swap(x1, x2)
	fmt.Printf("返回多个值的函数：原来是%s,%s, 交换位置后得到%s, %s\n", x1, x2, res1, res2)

	println()
	println()

	println("看看Go语言 string是否是值传递的：")
	mike :="Mike"
	playSwitch(mike)
	changePlayer(mike, "Louis")
	playSwitch(mike)
}

func changePlayer(name string, newName string) {
	println(name, "已经换人了，现在是", newName)
	name = newName
}

func playSwitch(player string) {
	println(player, "playing the Switch ~")
}

func swap(x1 string, x2 string) (string, string) {
	return x2, x1;
}

func calculateFinalSalary(base float64, promoteRate float32) float64 {
	return base * float64(promoteRate)
}

func add(a int, b int) int {
	return a + b
}

func printInt() int {
	return 1
}
