package main

import "fmt"

func ThirdWay() {
	var j, s string
	hub := map[rune]int{}
	var output int
	fmt.Scan(&j, &s)
	for _, v := range j {
		hub[v]++
	}
	for _, v := range s {
		if hub[v] >= 1 {
			output++
		}
	}
	fmt.Print(output)

}
