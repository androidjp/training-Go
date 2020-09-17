package main

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := GetConfig()
	fmt.Println(cfg.IsRotate())
	cfg.SetRotate(true)
	fmt.Println(cfg.IsRotate())
}
