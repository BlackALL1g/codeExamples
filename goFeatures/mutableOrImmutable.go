// Program to demonstrate triangle pic that you create with paticular parameters

package goFeatures

import (
	"fmt"
)

type first struct {
	field1 int
	field2 string
}

func MutableOrImmutable() {
	v := first{field1: 45, field2: "good"}
	// f.field1 := 45
	// f.field2("go")
	fmt.Println(v.field2)

	// You reassign var y
	x := 7
	y := x
	y = 9
	fmt.Println(x, y)

	// You just rename the same slice
	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := sl1
	sl2[0] = 0
	sl2[1] = 100
	sl2[2] = 200
	fmt.Println(sl1, sl2)

	// You just rename the same map
	mp1 := map[string]int{"good": 123, "day": 456}
	fmt.Println(mp1)
	mp2 := mp1
	mp2["good"] = 789000
	mp2["day"] = 789000
	fmt.Println(mp1, mp2)
}
