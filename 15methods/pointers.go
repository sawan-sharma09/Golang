package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (p Student) getAge() {
	fmt.Println(p.age)
}

func (s *Student) setAge() { //here the age will be changed by using pointers
	s.age = 25
	fmt.Println(s.age)

}

func (a Student) getAverageGrades() float32 {
	sum := 0
	for _, val := range a.grades {
		sum += val
	}

	return float32(sum) / float32(len(a.grades))
}

func main() {
	s1 := Student{"sawan", []int{12, 88, 79, 37, 28, 62}, 24}
	s2 := Student{"sawan", []int{12, 88, 79, 37, 28, 62, 76, 89, 90, 95}, 24}
	s1.getAge()
	fmt.Println(s1)
	s1.setAge()
	fmt.Println(s1)

	average := s1.getAverageGrades()
	fmt.Println(average)

	fmt.Println(s2.getAverageGrades()) //it works same as as the above 2 lines as we can directly store the return value and print it
}

/*package main    //This program seems better to me than the above one

import "fmt"

type Student struct {
	name  string
	marks []int
	age   int
}

func (p *Student) getName() string {
	fmt.Println(p.name)
	p.name = "Sheru"
	// fmt.Println(p.name)
	return p.name

}

func (g Student) getMarks() float32 {
	fmt.Println(g.marks)
	sum := 0
	for _, val := range g.marks {
		sum += val
	}
	return float32(sum) / float32(len(g.marks))
}

func main() {
	s1 := Student{"Sawan", []int{98, 78, 82, 85, 28, 92}, 24}
	s2 := Student{"Sheru", []int{98, 78, 82, 85, 28, 92, 56, 78, 99}, 6}

	// s1.getName()
	// fmt.Println(s1.age)
	fmt.Println(s1.getName())
	average := s1.getMarks()
	fmt.Println(average)

	s2.getName()
	fmt.Println(s2.getMarks())

}
*/
