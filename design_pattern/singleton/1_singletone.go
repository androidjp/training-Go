package singleton

import "sync"

// 1. 结构体 代替 类
type Tool struct {
	values int
}

// 2. 创建私有变量
var instance *Tool

var once sync.Once

// 3. once.Do 只会执行一次
func GetInstance() *Tool {
	once.Do(func() {
		instance = new(Tool)
	})
	return instance
}

func (tool *Tool) GetValues() int {
	return tool.values
}

func (tool *Tool) SetValues(x int) {
	tool.values = x
}
