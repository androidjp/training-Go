package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// 声明一个等待组
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	// 遍历这些地址
	for _, url := range urls {
		// 每个任务开始时，将等待组加1
		wg.Add(1)

		// 开启一个并发
		go func(url string) {
			fmt.Println("开始协程")
			// 使用defer，表示函数完成时将等待组 减1
			defer wg.Done()
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}

	// 等待所有任务完成
	wg.Wait()

	fmt.Println("over")
}
