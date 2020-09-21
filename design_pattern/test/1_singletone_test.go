package test

import (
	"../singleton"
	"fmt"
	"testing"
)

func TestTool(t *testing.T) {
	myTool := singleton.GetInstance()
	myTool.SetValues(5)
	fmt.Println(myTool.GetValues()) // 5
	myTool.SetValues(10)

	myTool2 := singleton.GetInstance()
	fmt.Println(myTool2.GetValues()) // 10
}
