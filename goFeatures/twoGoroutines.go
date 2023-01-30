// Create two goroutines so that the numbers from one channel are read as they come in, squared,
// and the result is written to the second channel.

package goFeatures

import (
	"fmt"
	"math"
)

func TwoGoroutines() {
	naturals := make(chan float64)
	squares := make(chan float64)

	go func() {
		// First way
		for x := 0.0; x <= 10; x++ {
			x1 := math.Pow(x, 2)
			naturals <- x1
		}
		// Second way
		// for x := 0.0; x <= 10; x++ {
		// 	naturals <- x * x
		// }
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
