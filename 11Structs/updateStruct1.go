package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func main() {
	demo2()
}

func demo1() {
	pt1 := Point{1, 2}
	pt2 := Point{-5, 8}
	pt1.x = -44 //changing the value of x
	pt2.y = 100
	fmt.Println(pt1)
	fmt.Println(pt2)
}

func demo2() { //         assigning a particular value at a time, or if we don't want to assign all the values.
	pt1 := Point{y: 10}
	fmt.Println(pt1)

	pt2 := Point{x: 9}
	fmt.Println(pt2)
}
