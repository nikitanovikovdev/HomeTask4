package main

import "fmt"

func main() {
	var a int = 1
	var b int = 3

	if a <= 1 && b >= 3 {
		fmt.Println("a + b = ", a + b)
	} else {
		fmt.Println("a - b = ", a - b)
	}
}
