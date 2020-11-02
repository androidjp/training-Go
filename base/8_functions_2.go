package main

import (
	"errors"
	"fmt"
)

type operation func(x, y int) int

func DoCalculation(x, y int, op operation) int {
	if op == nil {
		return 0
	}
	return op(x, y)
}

// 闭包函数类型
type calculateFunc func(x, y int) (int, error)

func GenCalculateFunc(op operation) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("eerrrrr")
		}
		return op(x, y), nil
	}
}

func main() {
	// 高阶函数： 入参为 函数， 或者 返回值 是 函数 的函数
	res := DoCalculation(1, 2, func(x, y int) int {
		return x + y
	})

	fmt.Println(res)

	// 闭包: 将自由对象，从不确定，变为确定
	var add = GenCalculateFunc(func(x, y int) int {
		return x + y
	})

	res, err := add(3, 4)
	if err == nil {
		println(res)
	}

}
