// package main

// import "fmt"

// type Details struct {
// 	Name, Address, City string

// 	Road, Pincode int
// }

// func main() {
// 	myDetails := Details{
// 		Name:    "Sawan",
// 		Address: "Kesora",
// 		City:    "BBSR",
// 	}

// 	fmt.Println(myDetails)
// }

package main

import "fmt"

type Details struct {
	Name, City  string
	Roll, Class int
}

func main() {
	User := Details{Name: "Sawan", City: "Bbsr", Roll: 10, Class: 2}
	fmt.Println(User)
	fmt.Println(User.City)
	fmt.Println(User.Name)
	fmt.Println(User.Roll)
	fmt.Printf("The details are %+v", User)
}
