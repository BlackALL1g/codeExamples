package goFeatures

import (
	"fmt"
	"sort"
)

// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }

func SortSliceOfIntegers() {

	// a := []int{1, 2, 3}
	// reverse(a)
	// fmt.Println(a)
	// // Output: [3 2 1]

	a := []int{1, 2, 3}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	// Output: [3 2 1]
}
