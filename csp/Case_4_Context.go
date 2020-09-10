package main

import (
	"context"
	"fmt"
	"time"
)

const DB_ADDRESS = "db_address"
const CALCULATE_VALUE = "calculate_value"

func readDB(ctx context.Context, cost time.Duration) {
	fmt.Println("db address is", ctx.Value(DB_ADDRESS), ", now start to connect DB.")
	select {
	case <-time.After(cost): // 模拟数据库读取
		fmt.Println("[FINISH] read data from db")
	case <-ctx.Done(): // 任务取消的原因
		fmt.Println("ctx.Done()...:", ctx.Err())
		// 一些清理工作
	}
}

func calculate(ctx context.Context, cost time.Duration) {
	fmt.Println("calculate value is", ctx.Value(CALCULATE_VALUE), ", now start calculate.")
	select {
	case <-time.After(cost):
		fmt.Println("[FINISH] calculate finish")
	case <-ctx.Done():
		fmt.Println("ctx.Done()...:", ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	// 添加上下文信息
	ctx = context.WithValue(ctx, DB_ADDRESS, "localhost:10086")
	ctx = context.WithValue(ctx, CALCULATE_VALUE, 1234)
	// 设定 子 Context 2s后执行超时返回
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	// 读DB 用时 4s
	go readDB(ctx, time.Second*4)
	// 计算 用时 4s
	go calculate(ctx, time.Second*4)
	// 主协程 充分执行
	time.Sleep(time.Second * 5)
}
