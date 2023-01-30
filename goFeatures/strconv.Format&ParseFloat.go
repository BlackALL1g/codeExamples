package goFeatures

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		// panic(e)
		fmt.Println(e)
	}
}

func StrconvFormatParseFloat() {
	var f = 65.46856
	n := strconv.FormatFloat(f, 'f', 4, 64)
	fmt.Println(n)

	s := "3.14159265"

	if f, err := strconv.ParseFloat(s, 32); err == nil {
		fmt.Println(f) // 3.1415927410125732
		check(err)
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		fmt.Println(f) // 3.14159265
	}
}
