package main

import "sync"

// 1. 结构体 代替 类
type Tool struct {
	values int
}

// 2. 创建私有变量
var instance *Tool

// 3. 锁对象（用于保证线程安全）
var lock sync.Mutex

// 4. 获取单例对象的方法，引用传递 返回
func GetInstance() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}

func (tool *Tool) GetValues() int {
	return tool.values
}

func (tool *Tool) SetValues(x int) {
	tool.values = x
}
