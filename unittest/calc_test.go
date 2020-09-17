package main

import (
	"fmt"
	"testing"
)

/**
cd到unittest目录，然后直接命令行运行： go test
如果想要显示详细的每个测试方法是否验证成功： go test -v
如果想指定只跑某个测试：go test -run TestAdd -v

t.Fatal/t.Fatalf 遇错即停
t.Error/t.Errorf 遇错不停
*/
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	if ans := Mul(1, 2); ans != 2 {
		t.Errorf("1 + 2 expected be 2, but %d got", ans)
	}

	if ans := Mul(-10, -20); ans != 200 {
		t.Errorf("-10 + -20 expected be 200, but %d got", ans)
	}
}

func TestMul_SubTests(t *testing.T) {
	t.Run("should_return_6_when_Mul_given_2_and_3", func(t * testing.T) {
		if ans := Mul(2, 3); ans != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("should_return_negative_6_when_Mul_given_2_and_negative_3", func(t * testing.T) {
		if ans := Mul(2, -3); ans != -6 {
			t.Fatal("fail")
		}
	})
}

func setup() {
	fmt.Println("[[[[[[[[[[ setup")
}

func teardown() {
	fmt.Println("[[[[[[[[[[ teardown")
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}