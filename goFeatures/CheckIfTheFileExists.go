package goFeatures

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("/home/vadim/Desktop/go_all_files/src/draft.go")
	// f := os.IsExist(err)
	// fmt.Println(f)
	// fmt.Println(err)

	if err != nil {
		fmt.Println("There is no such file")
		return
	}

	if err == nil {
		fmt.Println("Yep, i see this!")
		return
	}

	// file, err := os.Open("data.json")
	// if errors.Is(err, os.ErrNotExist) {
	// 	fmt.Println("File not found")
	// 	return
	// }

	// // Выполните какое-нибудь действие с файлом
	// defer file.Close()

}
