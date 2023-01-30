package goFeatures

import "fmt"

func returner2() (int, int, float64, string, bool) {
	return 5, 10, 76.9, "GOLD", false
}

func returner() {

	_, b, _, d, e := returner2()

	fmt.Println()
	fmt.Println(b)
	fmt.Println()
	fmt.Println(d)
	fmt.Println(e)
}
