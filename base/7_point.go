package main

import "fmt"

func main() {
	var x uint8 = 3
	var y uint8 = 2
	fmt.Println(x >> y)// 0
	// 获取x 的指针
	xStr := &x
	fmt.Println(xStr)  // 打印指针的值（也就是x值的内存地址）： 0xc00000a0b0
	fmt.Println(*xStr) // 3

	println()
	println()
	println("如何使用指针")
	var a int = 20
	var ip *int
	ip = &a
	println("a 值", a)
	println("a 变量地址", &a)
	println("ip 变量的值", ip)
	println("通过 （指针）*ip 返回值", *ip)

	println()
	println()
	println("空指针")
	var nilIp *int
	println("默认指针为空，值为", nilIp)
	println("到底是不是空指针的判断：", "if(nilIp == nil)", "结果是", nilIp == nil)

}
