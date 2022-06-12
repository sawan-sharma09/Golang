package main

import (
	"fmt"
	"math"
)

type areaPeri interface {
	area() float32
	perimeter() float32
}
type calArea interface {
	area() float32
}
type calPeri interface {
	perimeter() float32
}
type circle struct {
	radius float32
}
type rect struct {
	length, breath float32
}

func (a circle) area() float32 {
	circArea := math.Pi * a.radius * a.radius
	return circArea
}
func (p circle) perimeter() float32 {
	return 2 * math.Pi * p.radius
}
func (a rect) area() float32 {
	rectArea := a.length * a.breath
	return rectArea
}
func (p rect) perimeter() float32 {
	return 2 * (p.length * p.breath)
}
func main() {
	c := circle{10.5}
	r := rect{length: 4, breath: 5}

	arPer := []areaPeri{r, c}
	for _, val := range arPer {
		checkAreaPeri(val)
	}
}

func checkAreaPeri(ap areaPeri) {
	if val, ok := ap.(circle); ok {
		fmt.Println(val)
		fmt.Println("Area of circle is:", ap.area())
		fmt.Println("Perimeter of circle is:", ap.perimeter())
	} else if val1, ok1 := ap.(rect); ok1 {
		fmt.Println(val1)
		fmt.Println("area of the rectrangle is:", ap.area())
		fmt.Println("Perimeter of the rectrangle is:", ap.perimeter())
	} else {
		fmt.Println("Invalid record found")
	}
}
