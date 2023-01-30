// Go program to illustrate the
// use of Misc Operators
package otherStuff

import "fmt"

func PointersInAction() {
	a := 4

	// Using address of operator(&) and
	// pointer indirection(*) operator
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
