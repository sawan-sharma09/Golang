package main

import (
	"fmt"
	"math"
)

type volAreaPerimeter interface {
	volume() float64
	AreaPerimeter
}

type AreaPerimeter interface {
	calArea
	calPerimeter
}

type calPerimeter interface {
	perimeter() float64
}
type calArea interface {
	area() float64
}
type circle struct {
	radius float64
}

type rect struct {
	width, height, length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c1 circle) perimeter() float64 {
	return 1.1
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r1 rect) perimeter() float64 {
	return 2 * (r1.height + r1.width)
}
func (v rect) volume() float64 {
	return v.width * v.height * v.length
}
func main() {
	c1 := circle{29.3}
	r1 := rect{10, 33, 20}

	totArea(c1)
	totArea(r1)

	totPerim(r1)
	finalAreaPerim(c1)
	finalAreaPerim(r1)

	vol(r1)

}

func totArea(a calArea) {
	fmt.Println(a.area())
	// fmt.Println("----------------")
}

func totPerim(p calPerimeter) {
	fmt.Println("+++++++")
	fmt.Println(p.perimeter())
}
func finalAreaPerim(f AreaPerimeter) {
	fmt.Println(f.area())
	fmt.Println(f.perimeter())
}
func vol(v1 volAreaPerimeter) {
	fmt.Println("________________________")
	fmt.Println(v1.area())
	fmt.Println(v1.perimeter())
	fmt.Println(v1.volume())
}
