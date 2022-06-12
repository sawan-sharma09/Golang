package main

import "fmt"

type Marks struct {
	Physics, Chem, Maths int
}

type Student struct {
	Name         string
	Class, Roll  int
	SubjectMarks []Marks
}

func main() {
	details := Student{
		Name:  "Sheru",
		Class: 10,
		Roll:  001,
		SubjectMarks: []Marks{
			Marks{
				Physics: 20,
				Chem:    30,
				Maths:   58,
			},
			Marks{
				Physics: 77,
				Chem:    85,
				Maths:   22,
			},
			Marks{
				Physics: 20,
				Chem:    30,
				Maths:   58,
			},
		},
	}
	fmt.Println(details)
	fmt.Println("Details of Student: \n")

	fmt.Println("Student's name :", details.Name)
	fmt.Println("Student's class :", details.Class)
	fmt.Println("Student's Roll :", details.Roll)

	fmt.Println("Details of Marks: \n")
	fmt.Println(details.SubjectMarks)
	fmt.Println(details.SubjectMarks[2])

	for key := range details { /// to check if we can range over a struct
		fmt.Println(key)
	}
}
