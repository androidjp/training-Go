package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量
var (
	l  sync.Mutex
	v1 int
)

// 修改共享变量
func change(i int) {
	l.Lock()
	// ----- 临界区 start
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	// ----- 临界区 end
	l.Unlock()
}

// 访问共享变量
func read() int {
	l.Lock()
	a := v1 // 临界区
	l.Unlock()
	return a
}

// 保证了同一时刻只有一个协程能执行change or read
func main() {
	var numGR = 21 // 启动21个协程，
	var wg sync.WaitGroup
	fmt.Printf("%d", read())
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 每个协程 都 先修改 后读取，先后经历 两次Lock和两次UnLock
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}
	wg.Wait()
}

// main() 每次执行，得到的结果存在随机性，毕竟第i个协程是否在第i+1个协程之前执行，是不知道的
