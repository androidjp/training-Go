package main

import (
	"errors"
	"fmt"
	"math"
)

/**
Go 内置的 error类型 是一个接口类型，其定义如下：
type error interface {
  Error() string
}
*/

// 这是一个 利用了 error 接口的 方法
func mySqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	println("--------------------------")
	println("基础用法")
	println("--------------------------")
	result, err := mySqrt(-1)
	if err != nil {
		fmt.Printf("报错了：%v\n", err)
	} else {
		println("结果是：", result)
	}

	println("--------------------------")
	println("自定义实现 error 接口")
	println("--------------------------")

}
