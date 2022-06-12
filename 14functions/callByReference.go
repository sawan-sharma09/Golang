package main

import "fmt"

func main() {
	x, y := 12, 30
	fmt.Printf("Before swapping: x=%v and y=%v\n", x, y)
	swap(&x, &y)
	fmt.Printf("After Swapping: x=%v and y=%v", x, y) // the value changed because we passed the actual value using pointers.
}
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

//----------------------------------------------------------------------------

package main

import "fmt"

func cal(x, y *int) {
	*x = 55
	*y = 66

}

func main() {
	a, b := 20, 39
	temp := a
	a = b
	b = temp
	fmt.Printf("a=%v\n", a)
	fmt.Println("b=", b)
	// fmt.Println(temp)

	cal(&a, &b)
	fmt.Println(a, b)
}