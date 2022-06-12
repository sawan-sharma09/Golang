package main

import (
	"fmt"
	"math" //                     Use of Multiple interfaces
)

type findPerimeter interface {
	perimeter() float64
}

type findArea interface {
	area() float64
}

type findVolume interface {
	// volume() float64   /*I didn't use float64 here because there was no formula of volume for both circle and rectrangle*?
	volume() string
}

type circle struct {
	radius float64
}
type rectrangle struct {
	length, breath float64
}

func (c1 circle) area() float64 {
	return math.Pi * c1.radius * c1.radius
}

func (v1 circle) volume() string {
	return fmt.Sprint("There's no volume of Circle......")
}

func (r1 rectrangle) area() float64 {
	return r1.length * r1.breath
}

func (r2 rectrangle) perimeter() float64 {
	return 2 * (r2.length * r2.breath)
}

func (v2 rectrangle) volume() string {
	return fmt.Sprint("There is no volume of Rectrangle..")
}
func main() {
	c := circle{10.0}
	r := rectrangle{20, 44}

	letsCheckArea(c)
	letsCheckArea(r)

	letsCheckPerim(r)

	letsCheckVolume(c)
	letsCheckVolume(r)
}
func letsCheckArea(a findArea) {
	fmt.Println(a.area())
}
func letsCheckPerim(p findPerimeter) {
	fmt.Printf("Perimeter of rectrangle is %v\n", p.perimeter())
}

func letsCheckVolume(v findVolume) {
	fmt.Println(v.volume())
}
