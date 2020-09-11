package main

/**
逃逸分析命令：go tool compile -m main.go
显示程序的汇编代码：go tool compile -S main.go
详情看文章：https://juejin.im/post/6869969064501248008?utm_source=gold_browser_extension#heading-0
*/
type smallStruct struct {
	a, b int64
	c, d float64
}

//go:noinline
func main() {
	smallAllocation()
}

func smallAllocation() *smallStruct {
	return &smallStruct{}
}
