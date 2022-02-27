package main

import "fmt"

func main() {
	var nums = []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(i)
		//	functions like enumerate in python
		//	i will be indexes and num each number in the array
	}

	// for maps it's gonna be key and value

	fmt.Println(sum)
	maps := map[string]string{"a": "apple", "b": "banana"}

	for key, value := range maps {
		fmt.Println(key, value)
	}

	for key, _ := range maps {
		fmt.Println(key, maps[key])
	}

	for key := range maps {
		fmt.Println(key, maps[key])
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	} // by default, it's gonna iterate chars unicode characters

	str := "abc"

	individualChars := []rune(str)
	fmt.Println(individualChars) // still not usefull

	//for i := 0; i < len(individualChars); i++ {
	//	fmt.Println(string(individualChars[i]))
	//}

	for _, char := range individualChars {
		fmt.Println(string(char))
	}
}
