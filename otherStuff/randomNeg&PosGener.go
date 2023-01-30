package otherStuff

import (
	"fmt"
	"math/rand"
	"time"
)

// var fields = [2][2]int{
// 	{1,2,3}
// 	{4,5,6}
// }

// var a = [3][3][4]int{
// 	{
// 		{0, 1, 2, 3},
// 		{4, 5, 6, 7},
// 		{8, 9, 10, 11},
// 	},
// 	{
// 		{0, 1, 2, 3},
// 		{4, 5, 6, 7},
// 		{8, 9, 10, 11},
// 	},
// 	{
// 		{0, 1, 2, 3},
// 		{4, 5, 6, 7},
// 		{8, 9, 10, 11},
// 	},
// }

// func main() {
// 	/* an array with 5 rows and 2 columns*/
// 	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
// 	var i, j int

// 	/* output each array element's value */
// 	for i = 0; i < 5; i++ {
// 		for j = 0; j < 2; j++ {
// 			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
// 		}
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	lim := rand.Perm(11)

// 	arr := [5][5]int{}
// 	copy(arr, lim)
// 	fmt.Println(arr)

// 	// for i, v := range arr {
// 	// 	fmt.Println(v, i)
// 	// }
// 	fmt.Println(lim)
// 	len()
// }

func RandomNegativeAndPositiveNyberGenerator() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println(randInt(-100, 100))

}

func randInt(min, max int) (g int) {
	return min + rand.Intn(max-min)
}

// func main() {

// 	data, err := generateRandomBytes(17)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(data)
// }

// func generateRandomBytes(n int) ([]byte, error) {

// 	b := make([]byte, n)
// 	_, err := rand.Read(b)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return b, nil
// }
