package main

import "fmt"

type Student struct {
	Name, Section string
	Class, Roll   int
}

type College struct {
	Name, City, Branch string
	Year               int
	Details            Student //   Here student is a type (working/ behaving same as the date types int/float etc)
}

func main() {
	clg := College{
		Name:   "CVRaman",
		City:   "BBSR",
		Branch: "CSE",
		Year:   2019,
		Details: Student{
			Name:    "Sawan",
			Section: "A",
			Class:   16,
			Roll:    24,
		},
	}
	fmt.Println("Details of the College")
	fmt.Println("College name:", clg.Name)
	fmt.Println("City :", clg.City)
	fmt.Println("Branch :", clg.Branch)
	fmt.Println("Year :", clg.Year)

	fmt.Println("\nDetails of Student")
	fmt.Println("College name:", clg.Details.Name)
	fmt.Println("Section:", clg.Details.Section)
	fmt.Println("Class:", clg.Details.Class)
	fmt.Println("Roll:", clg.Details.Roll)

	fmt.Println(clg)
	fmt.Printf("\nThe whole details are %+v", clg)

}
