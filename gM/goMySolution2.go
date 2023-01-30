package gM

import (
	"fmt"
	"regexp"
)

func GoMySolution2() {
	aORb := regexp.MustCompile("a|b")

	matches := aORb.FindAllStringIndex("aabbccd", -1)
	fmt.Println(len(matches))
}
