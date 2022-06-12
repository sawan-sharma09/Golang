package main

import "fmt"

var c interface{}

type person interface{}

type admin struct {
}

func main() {
	fmt.Printf("type of c is %T c: %v\n", c, c)
	d := &c
	fmt.Printf("type of d is %T d: %v\n", d, d)

	// e := &person{}
	var e person
	fmt.Printf("type of e is %T e: %v\n", e, e)

}
