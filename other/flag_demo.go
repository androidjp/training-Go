package main

import (
	"flag"
	"fmt"
)

var name string

func init() {

	flag.StringVar(&name, "name", "everyone", "The greeting obj.")

}

/**
直接运行：go run flag_demo.go -name="tttt"
最终结果：hello, tttt

运行：go run flag_demo.go --help
最终结果：
  -name string
        The greeting obj. (default "everyone")

*/
func main() {
	flag.Parse()
	fmt.Printf("hello, %s \n", name)
}
