package codeExamples

import "fmt"

func GolangExperience1() {
	ch := make(chan string)
	go pushToChannel(ch)
	for val := range ch {
		fmt.Println(val)
	}
}
func pushToChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)
}
