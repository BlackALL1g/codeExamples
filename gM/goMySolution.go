package codeExamples

import (
	"fmt"
	"strings"
)

func GoMySolution() {
	a := strings.Count("aabbccd", "a")
	b := strings.Count("aabbccd", "b")
	result := (a + b)
	fmt.Println(result)

}
