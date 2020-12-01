package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Tick(time.Second)
	for {
		select {
		case <-t1:
			fmt.Println("run...")
		}
	}
}
