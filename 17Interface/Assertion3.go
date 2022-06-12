package main

import "fmt"

type ninjawar interface {
	attack()
	retreat()
}

type ninjaweapon interface {
	attack()
}
type ninjasword struct {
	owner string
}
type ninjastar struct {
	owner string
}

func (n ninjastar) attack() {
	fmt.Println("attck the Ninjastar")
}
func (n ninjasword) attack() {
	fmt.Println("attack the ninjasword")
}
func (n ninjasword) retreat() {
	fmt.Println("retreat the attack")
}
func (n ninjasword) over() {
	fmt.Println("over the attack")
}

func main() {
	weapons := []ninjaweapon{ninjasword{"Sawan"}, ninjastar{"Sheru"}}
	for _, i := range weapons {
		letsCheckNinja(i)

	}
	letsCheckWar(ninjasword{"Baba"})

}
func letsCheckNinja(let ninjaweapon) {
	let.attack()
	fmt.Printf("%v\n", let)

	if val, ok := let.(ninjasword); ok {
		val.retreat()
	} else {
		fmt.Println(val, ok)
	}
}
func letsCheckWar(war ninjawar) {
	war.attack()
	war.retreat()
	fmt.Println("-------------------")
	if val, ok := war.(ninjasword); ok {
		val.over()
	} else {
		fmt.Println(val, ok)
	}
}
