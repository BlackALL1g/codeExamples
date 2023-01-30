package gE

import "fmt"

func GolangExperience1() {
	ch := make(chan string)
	go PushToChannel(ch)
	for val := range ch {
		fmt.Println(val)
	}
}
func PushToChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)
}
