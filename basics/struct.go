package main

import "fmt"

type person struct {
	name string
	age  int // in case not provided will be taken as 0
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 43
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(newPerson("Yalchin"))
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
}

// Structs are mutable.
// An & prefix yields a pointer to the struct.
// Omitted fields will be zero-valued.
