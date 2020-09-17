package main

import (
	"fmt"
	"testing"
)

func TestTool(t *testing.T) {
	myTool := GetInstance()
	myTool.SetValues(5)
	fmt.Println(myTool.GetValues()) // 5
	myTool.SetValues(10)

	myTool2 := GetInstance()
	fmt.Println(myTool2.GetValues()) // 10
}
