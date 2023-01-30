package gM

import "fmt"

func sumDiffmul(a int, b int, c int) (int, int, int) {

	return (a + b), (a - b), a * c
}

func GoMySolution3() {

	var a = 68
	var b = 100
	var c = 50

	var sum, diff, mul = sumDiffmul(a, b, c)

	// Printing the values
	fmt.Println("Sum = ", sum, "\nDifference = ", diff, "\nMul=", mul)
}
