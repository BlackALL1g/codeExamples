package goFeatures

import "fmt"

func sum(s []int, mychannel chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	mychannel <- sum // send sum to channel
}

func ChannelLogis() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mychannel := make(chan int)
	go sum(s[:len(s)/5], mychannel) //3
	go sum(s[:], mychannel)         //55
	go sum(s[:len(s)/2], mychannel) //15
	go sum(s[:len(s)/3], mychannel) //6
	x := <-mychannel                // receive from mychannel
	y := <-mychannel                // receive from mychannel
	z := <-mychannel                // receive from mychannel
	j := <-mychannel                // receive from mychannel

	fmt.Println("Sum computed in first goroutine: ", x)
	fmt.Println("Sum computed in second goroutine: ", y)
	fmt.Println("Sum computed in third goroutine: ", z)
	fmt.Println("Sum computed in fourth goroutine: ", j)
	// fmt.Println("Total sum: ", x+y)
}
