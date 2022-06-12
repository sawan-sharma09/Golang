package main

import (
	"fmt"
	"math"
)

type AreaMaster interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (a1 Circle) area() float64 {
	circArea := math.Pi * a1.radius * a1.radius
	return circArea
}

type Rectrangle struct {
	length, breath float64
}

func (a2 Rectrangle) area() float64 {
	rectArea := a2.length * a2.breath
	return rectArea
}
func (p2 Rectrangle) perimeter() float64 {
	rectPerim := 2 * (p2.length + p2.breath)
	return rectPerim
}
func main() {
	c1 := Circle{10.2} //Here c1 is shown in read because the variable has not been used.
	r1 := Rectrangle{13, 15}
	// CirandRect(c1)
	/*cannot use c1 (variable of type Circle) as type AreaMaster in argument to CirandRect:
	Circle does not implement AreaMaster (missing perimeter method)*/
	CirandRect(r1)
}

func CirandRect(TotArea AreaMaster) {
	fmt.Println(TotArea.area())
	fmt.Println(TotArea.perimeter())
}
