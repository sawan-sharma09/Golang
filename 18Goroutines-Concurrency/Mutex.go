package main

import "fmt"

func main() {
	var (
		foo int
		bar string
	)
	fmt.Println("Please enter int")
	fmt.Scanln(&foo)
	fmt.Printf("foo: %v\n", foo)

	fmt.Println("Please enter string")
	fmt.Scanf("%v", &bar)
	fmt.Printf("bar: %v\n", bar)
}
