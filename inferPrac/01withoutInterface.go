package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (a1 Circle) area() float64 {
	cirArea := math.Pi * a1.radius * a1.radius
	return cirArea
}

type Rectrangle struct {
	length, breath float64
}

func (a2 Rectrangle) area() float64 {
	rectArea := a2.length * a2.breath
	return rectArea
}

func main() {
	c1 := Circle{10.2}
	r1 := Rectrangle{10, 20}

	forCircle(c1)
	forRect(r1)
}

func forCircle(c Circle) {
	fmt.Println(c.area())
}

func forRect(r Rectrangle) {
	fmt.Println(r.area())
}
