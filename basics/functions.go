package main

import "fmt"

func main() {
	fmt.Println(plus(5, 4))
}

func plus(a int, b int) int {
	sum := a + b
	return sum
}

func plus2(a, b int) int {
	sum := a + b
	return sum
}
