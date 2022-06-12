package main

import (
	"fmt"
	"math"
)

type AreaPerimeter interface {
	AreaMaster
	PeriMaster
}

type AreaMaster interface {
	area() float64
	// perimeter() float64
}

type PeriMaster interface {
	perimeter() float64
}
type Circle struct {
	radius float64
}

func (a1 Circle) area() float64 {
	circArea := math.Pi * a1.radius * a1.radius
	fmt.Print("The area of the Circle is : ")
	return circArea
}

type Rectrangle struct {
	length, breath float64
}

func (a2 Rectrangle) area() float64 {
	rectArea := a2.length * a2.breath
	fmt.Print("The area of the Rectrangle is :")
	return rectArea
}
func (p Rectrangle) perimeter() float64 {
	rectPerim := 2 * (p.length + p.breath)
	fmt.Print("The perimeter of the Rectrangle is :")
	return rectPerim
}

func main() {
	c1 := Circle{20.6}
	r1 := Rectrangle{20, 49}

	totArea(r1)
	totArea(c1)
	totPerim(r1)

	bestInterface(r1)
}

func totArea(a AreaMaster) {
	fmt.Println(a.area())
	// fmt.Println(a.perimeter())
}
func totPerim(peri PeriMaster) {
	fmt.Println(peri.perimeter())
}

func bestInterface(b AreaPerimeter) {
	fmt.Println("..........")
	fmt.Println(b.area())
	fmt.Println(b.perimeter())

}
