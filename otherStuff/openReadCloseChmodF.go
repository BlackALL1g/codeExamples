package otherStuff

import (
	"fmt"
	"io"
	"log"
	"os"
)

// "/home/vadim/Desktop/myFiles/Images/Фабула.docx"
// 0cf98c50-7880-4a2a-b2a0-6512b139e9d8.pdf

func openReadCloseChmodF() {

	// os.Chmod("/etc/shadow", 0664)

	// f, err := os.OpenFile("/home/vadim/Desktop/myFiles/Images/Фабула.docx", os.O_RDWR|os.O_CREATE, 0764)

	f, err := os.Open("/home/vadim/Desktop/myFiles/Images/Фабула.docx")

	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := io.ReadAll(f)

	if err != nil {
		log.Fatal()
	}

	fmt.Println(n[0:10])

	// os.Chmod("/home/vadim/Desktop/myFiles/Images/Фабула.docx", 0400)

	// f.Chmod(0440)

	f.Close()

}
