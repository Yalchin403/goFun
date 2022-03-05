package main

import (
	"fmt"
	"github.com/Yalchin403/goFun/basics/mypck"
	"math"
)

type geometry interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func measure(g geometry) float64 {
	area := g.Area()
	return area
}

func main() {
	r := mypck.Rect{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	rArea := measure(r)
	cArea := measure(c)
	fmt.Println(cArea, rArea)
}
