package main

import "fmt"

type Person struct {
	name     string
	phNo     int
	location string
}

func (p Person) getPer() {
	fmt.Printf("%+v", p.name)
	fmt.Println(p.phNo)
	fmt.Println(p.location)
	p.name = "Sheru"

}
func main() {
	per1 := Person{"Sawan", 243525266, "BBSR"}

	fmt.Println(per1)
	per1.getPer()
	fmt.Println(per1)

}
