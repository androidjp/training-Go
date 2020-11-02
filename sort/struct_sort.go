package main

import (
	"fmt"
	"sort"
)

type Log struct {
	UserID     int
	Message    string
	Num        float64
	CreateTime string
}

type Wrapper struct {
	log []Log
	by  func(p, q *Log) bool
}

type SortBy func(p, q *Log) bool

func (pw Wrapper) Len() int { // 重写 Len() 方法
	return len(pw.log)
}
func (pw Wrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.log[i], pw.log[j] = pw.log[j], pw.log[i]
}
func (pw Wrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.log[i], &pw.log[j])
}

// 封装成 SortLog 方法
func SortLog(log []Log, by SortBy) {
	sort.Sort(Wrapper{log, by})
}

func main() {
	log := []Log{
		{1, "签到", 1, "1563935120"},
		{1, "充值", 100, "1563935320"},
	}

	fmt.Println(log)
	//调用wrapper
	sort.Sort(Wrapper{log, func(p, q *Log) bool {
		return q.Num < p.Num // Num 递减排序
	}})

	fmt.Println(log)
	//间接封装
	SortLog(log, func(p, q *Log) bool {
		return p.CreateTime < q.CreateTime // CreateTime 递增排序
	})

	fmt.Println(log)

}
