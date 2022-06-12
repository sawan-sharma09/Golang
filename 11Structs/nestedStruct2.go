package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	Name, Email, Address string
	Age                  int
	MonthlySalary        []Salary
}

func main() {
	emp := Employee{
		Name:    "Sawan",
		Email:   "Sksharma@gmail.com",
		Address: "kesora",
		Age:     24,
		MonthlySalary: []Salary{
			Salary{
				Basic: 10000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 20000.00,
				HRA:   5500.00,
				TA:    2700.00,
			},
			Salary{
				Basic: 30000.00,
				HRA:   6000.00,
				TA:    4000.00,
			},
		},
	}

	fmt.Println(emp)
	fmt.Println("\nDetails of employee are:\n")
	fmt.Println("Name is:", emp.Name)
	fmt.Println("Email is:", emp.Email)
	fmt.Println("Address is:", emp.Address)
	fmt.Println("Age is:", emp.Age)

	fmt.Println("Details of Salary are: \n")
	fmt.Println(emp.MonthlySalary)
	fmt.Println(emp.MonthlySalary[0])
	fmt.Println(emp.MonthlySalary[2])

	sal1 := emp.MonthlySalary[0]
	newSal := emp.MonthlySalary.Salary[0] // I want to change one of the salaries, but not able to do so
	fmt.Println(emp.MonthlySalary)
	fmt.Println(sal1)

}
