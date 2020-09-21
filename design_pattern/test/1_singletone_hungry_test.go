package test

import (
	"../singleton"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := singleton.GetConfig()
	fmt.Println(cfg.IsRotate())
	cfg.SetRotate(true)
	fmt.Println(cfg.IsRotate())
}
