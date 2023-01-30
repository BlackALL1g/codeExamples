package goFeatures

import "fmt"

func ReturnMultipleFunctionValues2() {
	var age, name = ReturnMultipleFunctionValues(27, 56, "Иван", "Иванов")
	fmt.Println(age)  // 27
	fmt.Println(name) // Иван Иванов
}

func ReturnMultipleFunctionValues(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}
