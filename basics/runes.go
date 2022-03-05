package main

import "fmt"

func main() {
	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	// unicode representation

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

}

/*
A range loop handles strings specially and decodes each rune along with its offset in the string.
We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
*/
