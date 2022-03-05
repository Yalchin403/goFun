package mypck

import (
	"fmt"
)

type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func main() {
	instance := Rect{5, 4}
	fmt.Println("area: ", instance.Area())
}
