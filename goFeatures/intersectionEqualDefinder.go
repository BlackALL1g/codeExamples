// f Equal, Intersection, Datatype Definder
package goFeatures

import (
	"fmt"
	"strings"
	"unicode"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}

	// I tired to iplement 	a = []int{1, 1, 1}
	// 						b = []int{1, 1, 1, 1}
	// 						Intersection(a, b)
	//						only output "1" instead of 111
	// for _, item := range b {
	// 	if _, ok := m[item]; ok {
	// 		c = append(c, item)
	// 	}
	// }

	return

}

// Data type definder

func Definder(x any) {
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)
}

// Absolutely the same, but remember that (any == interface{}) (true)
func Definder2(x interface{}) {
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)
}

// func CutOfLetters(s string) int {
// 	n, err := s.Read(b[byte])

//		return
//	}

func CutOfletters(a string) (n int) {
	a = strings.TrimFunc(a, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	return n
}

func Cat(a, b []int) {
	c := append(a, b...)
	for _, v := range c {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	return
}

func IntersectionEqualDefinder() {

	Definder("654.34")

	a := []int{37, 5, 1, 2}
	b := []int{6, 2, 4, 37}
	Intersection(a, b)
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	Intersection(a, b)
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	fmt.Println(Equal(a, b))

	a = []int{1, 2, 3}
	b = []int{4, 5, 6}
	// c := append(a, b...)
	// for _, v := range c {
	// 	fmt.Printf("%v ", v)
	// }
	// fmt.Println()
	Cat(a, b)

	a = []int{54, 345, 654}
	b = []int{76, 34, 643}

	Cat(a, b)

	h := "mfgah134517095aldrfgvh8h"
	n := CutOfletters(h)
	fmt.Println(n)
}
