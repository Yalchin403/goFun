package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a, b int = 1, 2 // initialize as int can't change later as it is statically typed
	var c, d = 3, 4     // initialize without stating type still not changeable
	e := 5

	const k = "API_KEY" // const variable goes here
	fmt.Println(a, b, c, d, e)
	fmt.Println(float64(123) / 12)
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).Kind().String()
}
