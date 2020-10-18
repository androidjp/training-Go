package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {

}

/**
自定义 help 的描述信息
运行：go run flag_demo.go --help
最终结果：
Usage of question:
  -name string
        The greeting obj. (default "everyone")



*/
func main() {
	var name = flag.String("name", "everyone", "The greeting obj.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("hello, %s \n", *name)
}
