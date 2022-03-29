package main

import "fmt"

/*
Let's say you need to make a function that takes one slice and prints it.
Simple, right? What if we want to have the slice be an integer? You will need to make a new method for that
generics are here to solve such problems
*/

//func main() {
//	m := map[int]int{1: 2, 2: 3, 3: 4}
//	fmt.Println("Keys m: ", MapKeys(m))
//}

//func MapKeys[K comparable, V any](m map[K]V) []K {
//	r := make([]K, 0, len(m))
//
//	for k := range m {
//		r = append(r, k)
//	}
//
//	return r
//
//}

// interface here means anytype

func Print(output ...interface{}) {
	fmt.Println(output)
}

//we can make another function by using generics
// this will force us to be consistent with the data we pass

func Print2[T interface{}](output ...T) {
	fmt.Println(output)
}

// we can have our own custom types too

type Type interface {
	int | string
}

// we can also create generic structs

type personClass[intStr any] struct {
	name intStr
	age  int
}

// create and instance of the generic struct

func Print3[T Type](output ...T) {
	fmt.Println(output)
}
func main() {
	Print(1, 2, 3)
	Print("1", 2, 3.0) // not statically typed
	Print2(1, 2, 4)
	// if we now try to pass incosistent data
	//Print2(1, "2", 4.0)  // error
	Print3(1, 2, 77)
	// use generic personClass struct

	// make sure you instantiate the struct first just like that:node := Node[int]{}
	p := personClass[string]{name: "Yalchin403", age: 19}
	p2 := personClass[int]{name: 403, age: 24}
	fmt.Println(p, p2)
}
