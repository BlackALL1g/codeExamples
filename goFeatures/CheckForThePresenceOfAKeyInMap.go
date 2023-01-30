package goFeatures

import "fmt"

// func CheckTheKeyExisting(m map[any]any) bool {

// 	for k := range m {
// 		if k == x {
// 			return false
// 		}
// 	}

// 	return true
// }

func CheckTheKeyExisting() {

	mp := map[string]int{
		"six":  6,
		"five": 5,
	}

	k := mp["five"] == 4
	fmt.Println(k)

	brandsMap := map[string]string{
		"ford": "ford",
		"audi": "ford",
		"lada": "nil",
	}
	_, value := brandsMap["lada"]
	fmt.Println(value) // true,
	//value выдаст true или false, если ключ lada есть в brandsMap

	// a := map[int]string{5: "five", 6: "six"}
	// fmt.Println(CheckTheKeyExisting(a[4]))

}
