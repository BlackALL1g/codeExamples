// Rocks and jewels
// You have two lines of lowercase Latin characters: line J and line S.
// The characters in the line J are "jewels", in line S are "stones".
// You need to determine how many "stones" are "jewels" at the same time.
// Put simply, you need to check how many characters from S are also in J.
// This is a warm-up task which have ready-made solutions.
// It is very simple and is necessary to get comfortable with Yandex.Contest.
// Input and output can be done through ﬁles, or through standard I/O streams, as you prefer.

package main

import (
	"fmt"
)

func main() {
	var s, j string

	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &j)

	mp := map[rune]struct{}{}
	for _, l := range s {
		mp[l] = struct{}{}
	}

	result := 0
	for _, v := range j {
		if _, value := mp[v]; value {
			result++
		}
	}

	fmt.Println(result)
}
