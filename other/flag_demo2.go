package main

import (
	"flag"
	"fmt"
)

func init() {

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
	var name = flag.String("name", "everyone", "The greeting obj.")
	flag.Parse()
	fmt.Printf("hello, %s \n", *name)
}
