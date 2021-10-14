package main

/**
go
*/

import (
	"fmt"
	"math"
)

type A struct {
	width  float64
	height float64
}

type B struct {
	radius float64
}

type Areable interface {
	area()
}

func (a A) area() float64 {
	return a.width * a.height
}

func (B B) area() float64 {
	return B.radius * B.radius * math.Pi
}

func main() {

	a := A{10, 23.6}
	b := B{6.4}
	fmt.Println(a.area())
	fmt.Println(b.area())
}
