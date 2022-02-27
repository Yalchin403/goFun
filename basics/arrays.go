package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("empty:", arr)

	for i := 0; i < 5; i++ {
		arr[i] = i
	}

	fmt.Println(arr)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
