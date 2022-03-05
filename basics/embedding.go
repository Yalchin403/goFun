package main

import "fmt"

type cat struct {
	name string
}

func (c cat) meow() {
	fmt.Println("Meow, Meow...")
}

type catChild struct {
	cat
	childName string
}

func (c cat) eat() {
	fmt.Println("Cat child eats")
}

func main() {
	childInstance := catChild{
		cat:       cat{},
		childName: "Mesi",
	}
	fmt.Println(childInstance.childName)
	childInstance.eat()
	childInstance.meow()
}
