package codeExamples

import "fmt"

// a slice variable with type int declared
type volleyball []float64

func GoMySolution7() {
	//an instance of mySlice initialized using short variable declaration

	new_variable := new(volleyball)

	//adding data to slice1 by first creating a temporary variable

	one_more_variable := append(*new_variable, 1.5, 2.2, 3.3, 4.4)
	new_variable = &one_more_variable
	fmt.Println(new_variable)
}
