package gE

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanExperience() {
	for {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("It's not an int! type integer number")
				break
			}
			// fmt.Println(f)
			fmt.Printf("%.1d\n", input)
			// var fl float64 = 45.56543
			// st := strconv.FormatFloat(fl, 'f', 1, 64)
			// fmt.Println(st)
		}
	}
}
