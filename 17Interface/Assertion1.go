package main

import "fmt"

type ninjaweapon interface {
	attack()
}

type ninjastar struct {
	owner string
}
type ninjasword struct {
	owner string
}

func (ninjastar) attack() {
	fmt.Println("throwing ninja star")
}
func (ninjastar) pickup() {
	fmt.Println("Picking up the ninja star")
}
func (ninjasword) attack() {
	fmt.Println("thrrowing ninja sword")
}
func main() {
	weapons := []ninjaweapon{ninjastar{"Sawan"}, ninjasword{"Sheru"}}
	for _, i := range weapons {
		letsCheckNinja(i)
	}

}
func letsCheckNinja(l ninjaweapon) {
	l.attack()
	val, ok := l.(ninjastar)
	if ok {
		val.pickup()
	}
}
