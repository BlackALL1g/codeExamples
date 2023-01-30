package main

import (
	"bufio"
	"fmt"
	"os"
)

func secondWay() {
	// var s, j string
	stones := bufio.NewScanner(os.Stdin)
	stones.Scan()
	jewels := bufio.NewScanner(os.Stdin)
	jewels.Scan()
	s := stones.Text()
	j := jewels.Text()

	// fmt.Println(strings.Compare(s, j))

	mp := map[rune]struct{}{}
	for _, v := range s {
		mp[v] = struct{}{}
	}

	result := 0
	for _, v := range j {
		if _, stones := mp[v]; stones {
			result++
		}

	}
	fmt.Println(result)
}
