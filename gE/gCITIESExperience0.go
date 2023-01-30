package gE

import "fmt"

type Cities struct {
	name       [2]string
	rollNumber [2]float64
}

func GolangExperience0() {

	// create empty slice of struct pointers
	cities := []*Cities{}

	//create struct and append it to the slice
	ct := new(Cities)
	ct.name[0] = "London"
	ct.name[1] = "Moscow"
	ct.rollNumber[0] = 8.8
	ct.rollNumber[1] = 9.9
	cities = append(cities, ct)

	//create another struct
	ct = new(Cities)
	ct.name[0] = "Berlin"
	ct.name[1] = "Washington"
	ct.rollNumber[0] = 4.4
	ct.rollNumber[1] = 6.6
	cities = append(cities, ct)

	for c := range cities {
		n := cities[c]
		fmt.Println("Cities", *n)
	}
}
