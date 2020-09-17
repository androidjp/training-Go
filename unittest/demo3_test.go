package main

import "testing"

type myCase struct {
	Str string
	Expected string
}

// 帮助函数：用于将某些 common 的 code，重构出来
func createFirstLetterToUpperCase(t *testing.T, c *myCase) {
	// to.Helper() 用于在 运行  go test 时能够打印出报错对应的行号
	t.Helper()
	if ans := FirstLetterToUpperCase(c.Str); ans != c.Expected {
		t.Errorf("input is `%s`, expect output is `%s`, but actually output `%s`", c.Str, c.Expected, ans)
	}
}

func TestFirstLetterToUpperCase(t *testing.T) {
	createFirstLetterToUpperCase(t, &myCase{"hello", "Hello"})
	createFirstLetterToUpperCase(t, &myCase{"ok", "Ok"})
	createFirstLetterToUpperCase(t, &myCase{"Good", "Good"})
	createFirstLetterToUpperCase(t, &myCase{"GOOD", "Good"})
}