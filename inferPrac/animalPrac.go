package main

import "fmt"

type Human interface {
	talking() string
	Animals
	// walking() string    /*We can write in this way also, it will give the same result, as interface Animal consists of */
	// eating() string     /* same behaviours i.e walking and eating*/
}
type Animals interface { /*1st interface */
	eating() string
	walking() string
}
type man struct {
	name string
	age  int
}

func (man1 man) talking() string {
	return fmt.Sprintf("%v can talk and his age is %v", man1.name, man1.age)
}
func (man2 man) walking() string {
	return fmt.Sprintf("%v is walking", man2.name)
}
func (man3 man) eating() string {
	return fmt.Sprintf("%v is eating..", man3.name)
}

type dog struct {
	name, breed string
}

func (d1 dog) eating() string {
	return fmt.Sprintf("%v is able to eat and his breed is %v", d1.name, d1.breed)
}
func (d1 dog) walking() string {
	return fmt.Sprintf("%v is walking", d1.name)
}

func main() {
	m := man{
		name: "Sawan",
		age:  24,
	}
	d := dog{
		name:  "Sheru",
		breed: "labrador",
	}
	fmt.Println(d)

	letsCheckHuman(m)
	letsCheckAnimal(m)
	letsCheckAnimal(d)

	//letsCheck(d)
	/*cannot use d (variable of type dog) as type Human in argument to letsCheck:
	dog does not implement Human (missing talking method)*/
}

func letsCheckHuman(h Human) {
	defer fmt.Println(h.walking())
	defer fmt.Println(h.eating())
	defer fmt.Println(h.walking())
}
func letsCheckAnimal(anim Animals) {
	fmt.Println("--------------------")
	fmt.Println(anim.walking())
	fmt.Println(anim.eating())
}
