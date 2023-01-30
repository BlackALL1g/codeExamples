package otherStuff

import "fmt"

func Onemoretest1() {
	x := 0
	terms := []int{1, 2, 3, 4, 5, 6}
	for _, term := range terms {
		if term%2 == 0 {
			x += term
			fmt.Println(term)
		}
	}
}
