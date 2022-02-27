package main

import "fmt"

func add(a, b, c, d int) (int, int) {
	res1 := a + b
	res2 := c + d

	return res1, res2
}

func main() {
	fmt.Println(add(1, 2, 3, 4))
}
