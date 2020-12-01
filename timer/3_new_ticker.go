package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("run...")
			ticker.Reset(2 * time.Second)
			ticker.Stop()
		}
	}
}
