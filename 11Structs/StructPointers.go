// package main

// import "fmt"

// type Point struct {
// 	x int32
// 	y int32
// }

// func changeX(p *Point) {
// 	p.x = 200
// }

// func main() {
// 	pt1 := &Point{-2, 5}
// 	fmt.Println(pt1)
// 	changeX(pt1)
// 	fmt.Println(pt1)
// }

package main

import "fmt"

type Point struct {
	x      int32
	y      float32
	z      bool
	origin string
}

func changeOri(p *Point) {
	p.origin = "Square"

}

func changeY(p *Point) {
	p.y = 97
}

func changeXandZ(p *Point) {
	p.x = 200
	p.z = false
}

func main() {
	pt1 := &Point{10, 23.4, true, "Circle"}
	fmt.Println(pt1)
	fmt.Println(pt1.origin)
	changeOri(pt1) //     changing the value of origin by passing a pointer to the changeOri functn
	fmt.Println(pt1.origin)

	changeY(pt1) //     changing the value of y by passing a Pointer to the ChangeY functn
	fmt.Println(pt1.y)

	changeXandZ(pt1) //       changing the value of x and z using Pointers
	fmt.Printf("value of x: %v, value of z: %v", pt1.x, pt1.z)
}
