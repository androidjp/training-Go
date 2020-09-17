package main

/**
benchmark 基准测试
*/
func Calculate100000Operation() int {
	x := 0
	for i := 0; i < 100000; i++ {
		x += i
	}
	return x
}
