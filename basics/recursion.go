package main

import "fmt"

var (
	cache = make(map[int]int)
)

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {

	//fmt.Println(fact(5))
	//fmt.Println(fibonacci(100)) // it is always software that makes a fast machine slow
	// test the speed
	//for i := 0; i < 40; i++ {
	//	fmt.Println(fibonacci(i))
	//}
	fmt.Println(fibonacciCached(100)) // fast enough? Sometimes it is not about the language, it's about algorithm

}

// favorite fibonacci

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func fibonacciCached(n int) int {
	_, prs := cache[n]

	if prs {
		return cache[n]
	} else {
		if n == 1 {
			return 1
		} else if n == 2 {
			return 1
		}

		result := fibonacciCached(n-2) + fibonacciCached(n-1)
		cache[n] = result

		return result
	}

}
