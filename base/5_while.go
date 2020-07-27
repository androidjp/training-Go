package main

import "fmt"

func main() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Println(a)
		a++
	}

	for i := 0; i < 10; i++ {
		println("i = ", i)
		if i == 5 {
			break
		}
	}
}
