package main

import "fmt"

func var_func(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	var_func(1, 2, 3, 4, 99)

	//	we can also use ... reverse as in js
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	res := var_func(nums...)
	fmt.Println(res)
}
