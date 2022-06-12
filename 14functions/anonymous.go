package main

import "fmt"

/*func main() {
	f := func(y string) int {
		fmt.Println("Hello", y)
		return 0

	}
	a := f("sk")
	fmt.Printf("Type is %T", f)
	fmt.Printf("Type is %T", a)

}*/

//------------------------------------------------------------------------------------------------------

func main() {
	a, b := 20, 33
	f := func(x int) (val1 string) {
		fmt.Println(a, b)
		a = 23
		b = 34
		fmt.Println(a, b)
		fmt.Println(x)
		val1 = "sawan"

		return

	}
	ret := f(10)
	fmt.Println(a, b)
	fmt.Printf("Type=%T and value of f is %v\n", f, f)
	fmt.Printf("Type =%T and value of ret is %v", ret, ret)

}
