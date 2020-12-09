package main

import "fmt"

type user struct {
	reg_no   int
	name     string
	subjects []string
}

func main() {
	var a1 = user{
		reg_no: 1,
		name:   "Arun",
		subjects: []string{
			"Eng",
			"Lang",
			"Social"},
	}

	var a2 = user{
		reg_no: 2,
		name:   "Ben",
		subjects: []string{
			"Eng",
			"Phy",
			"Maths"},
	}

	var a3 = user{
		reg_no: 3,
		name:   "Charles",
		subjects: []string{
			"Maths",
			"Computer_Science",
			"Biology"},
	}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	var a4 = a2
	a4.reg_no = 4
	a4.name = "faruq"
	fmt.Println(a4)
	fmt.Println("Excellent Student:", a1.name)
}
