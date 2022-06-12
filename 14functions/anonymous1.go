package main

import "fmt" //     Passing anonymous function as argument to another function

func test1(p func(int) int) {

	fmt.Println(p(10)) /* here the p(10) is written so that we can print the return value otherwise
	we would have only got 11,34 as answer and not the return value i.e: 70 */

}
func main() {
	test := func(x int) int {
		a, b := 11, 34
		fmt.Println(a, b)
		return x * 7
	}
	test1(test)
	/* test2 := func(x int) int {
		return x * 20
	 }

	test1(test2)*/

}

// --------------------------
// package main

// import "fmt"

// func temp(a func(m, n int) float32) {
// 	fmt.Println(a(3, 5))
// }
// func main() {
// 	f := func(x, y int) float32 {
// 		return float32(x * y)
// 	}
// 	temp(f)
// }
