package main

import (
	"fmt"
	"sync"
)

func input(ch chan string) {
	defer wg.Done()
	defer close(ch) // 当方法返回，就调用 ch关闭
	var input string
	fmt.Println("Enter 'EOF' to shut down:")
	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("Read input err: ", err.Error())
			break
		}
		if input == "EOF" {
			fmt.Println("Bye!")
			break
		}
		ch <- input
	}
}

func output(ch chan string) {
	// defer关键字后面跟一个函数：return 和 defer不是同时执行，而是 在return更新完返回值之后再去执行defer
	defer wg.Done()
	for value := range ch {
		fmt.Println("Your input: ", value)
	}
}

var wg sync.WaitGroup

func main() {
	/**
	从命令行 接受用户的输入并输出到命令行中
	*/
	ch := make(chan string)
	wg.Add(2)
	go input(ch)  // 协程一：读取输入
	go output(ch) // 协程二：输出到命令行。 协程间 通过 chan 进行数据传输，chan 中的数据传输遵循先进先出顺序，保证每次只能有一个协程发送或接收数据
	wg.Wait()
	fmt.Println("Exit!")
}
