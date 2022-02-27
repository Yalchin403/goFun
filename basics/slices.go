package main

import "fmt"

func main() {
	s := make([]string, 3) // []
	result := append(s, "random_str")
	result2 := append(result, "coder403")
	result3 := append(result2, "smth")
	result4 := append(result3, "smth2")
	fmt.Println("emp:", s, result, result2, result3, result4)
}
