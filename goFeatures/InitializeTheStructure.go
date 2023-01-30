//

package goFeatures

import "fmt"

type life struct {
	month string
	day   float64
}

type Student struct {
	Name string
	Age  int
}

func InitializeTheStructure() {

	var pa *Student   // pa == nil
	pa = new(Student) // pa == &Student{"", 0}
	pa.Name = "Alice" // pa == &Student{"Alice", 0}
	fmt.Println(pa)

	var g life
	g.month = "January"
	g.day = 4

	fmt.Println(g.month, g.day)

	// var way *life
	// way = new(life)
	way := life{
		month: "October",
		day:   6,
	}

	fmt.Println(way)

	superway := &life{
		month: "February",
		day:   19,
	}

	fmt.Println(superway.month, superway.day)

	c := life{"March", 8}

	fmt.Println(c)

	d := life{}

	fmt.Println(d)
}
