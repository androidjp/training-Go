package main

import "testing"

func TestMergeString(t *testing.T) {
	tests := []struct {
		name string
		X, Y, Expected string
	}{
		{"should_return_HelloWorld_when_MergeString_given_Hello_and_World", "Hello", "World", "HelloWorld"},
		{"should_return_aaaBBB_when_MergeString_given_aaa_and_BBB", "aaa", "bbb", "aaabbb"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ans:= MergeString(test.X, test.Y);ans!=test.Expected {
				t.Errorf("should return %s, but actually return %s", test.Expected, ans)
			}
		})
	}
}

