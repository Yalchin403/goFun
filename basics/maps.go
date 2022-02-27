package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8
	fmt.Println(m)
	fmt.Println(m["k2"])
	fmt.Println(len(m))

	_, prs := m["k3"]
	// underscore is blank identifier

	// if present, prs will be true
	fmt.Println("prs:", prs)

	//	initialize and declare at the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
