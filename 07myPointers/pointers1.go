package main

import "fmt"

func main() {
	demo2()
}

func demo1() {
	i, j := 43, 445
	fmt.Println(i, j)
	fmt.Println(&i, &j)

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	*p = *p + 2
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
}

func demo2() {
	i, j := 20, 60
	p := &i

	fmt.Println(p)
	fmt.Printf("%T\n", p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 30
	fmt.Println(j)
}

func demo3() {
	x := 7
	fmt.Println(&x)

	fmt.Println(*&x) //my experimentation
}

func demo4() {
	x := 4
	p := &x
	fmt.Println(x)
	*p = 10 //dereferencing the previous value of x and again putting  a new value in it.
	fmt.Println(*p)
	fmt.Println(x)
}
