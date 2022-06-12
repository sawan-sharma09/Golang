package main

import "fmt" /**Simple interface with only one behaviour i.e circle(),
and both vegetable and fruits implement this behaviour implicitly*/

type sameColour interface {
	circle() string
}
type vegetable struct {
	name, shape string
}
type fruits struct {
	name, shape string
}

func (v1 vegetable) circle() string {
	return fmt.Sprintf("The %v is of shape %v", v1.name, v1.shape)
}

func (f1 fruits) circle() string {
	return fmt.Sprintf("The %v is of shape %v", f1.name, f1.shape)
}
func main() {
	v := vegetable{name: "cabbage", shape: "circle"}
	f := fruits{name: "grapes", shape: "green"}

	letsCheckColour(v)
	letsCheckColour(f)
}

func letsCheckColour(clr sameColour) {
	fmt.Println(clr.circle())
}
