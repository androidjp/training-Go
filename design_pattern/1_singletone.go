package main

import "sync"

// 消息池
type messagePool struct {
	pool *sync.Pool
}

// 消息体
type Message struct {
	Count int
}

// 消息池单例
var msgPool = &messagePool{
	pool: &sync.Pool{New: func() interface{} {
		return &Message{Count: 0}
	}},
}

// 访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

// 添加消息
func (mp *messagePool) AddMsg(msg *Message) {
	mp.pool.Put(msg)
}

// 获取消息
func (mp *messagePool) GetMsg() *Message {
	return mp.pool.Get().(*Message)
}
