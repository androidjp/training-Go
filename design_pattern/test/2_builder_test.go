package test

import (
	"../builder"
	"testing"
)

func TestMessageBuilder(t *testing.T) {
	message := builder.Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.1").
		WithDestPort(1234).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Builder()
	if message.Header.SrcAddr != "192.168.0.1" {
		t.Errorf("expect src address 192.168.0.1, but actual %s.", message.Header.SrcAddr)
	}
	if message.Body.Items[0] != "record1" {
		t.Errorf("expect body item0 record1, but actual %s.", message.Body.Items[0])
	}
}
