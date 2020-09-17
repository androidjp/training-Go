package main

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkCalculate100000Operation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Calculate100000Operation()
	}
}

func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}