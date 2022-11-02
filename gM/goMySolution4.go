package codeExamples

import "fmt"

// struct definition
type Planes struct {
	name       string
	rollNumber int
}

func (g Planes) PrintDetails() {
	fmt.Printf("Planes Details\n---------------\n")
	fmt.Printf("Name : %s\n", g.name)
	fmt.Printf("Roll Number : %d\n", g.rollNumber)
}
func (k Planes) qwerty() {
	fmt.Printf("\nStudent Details\n---------------\n")
	fmt.Printf("Name : %s\n", k.name)
	fmt.Printf("Roll Number : %d\n", k.rollNumber)
}

func GoMySolution4() {
	var Good = Planes{name: "Anu", rollNumber: 14}
	Good.PrintDetails()
	var Bad = Planes{name: "erty", rollNumber: 16}
	Bad.qwerty()

}
