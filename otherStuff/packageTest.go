package otherStuff

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func check(e error) {
// 	if e != nil {
// 		log.Fatal()
// 	}
// }

// "/home/vadim/Desktop/myFiles/forGoTest.odt"

func PackageTest() {
	// dat, err := os.ReadFile("/home/vadim/Desktop/myFiles/forGoTest.odt")
	// check(err)
	// fmt.Print(string(dat[0:20]))

	// f, err := os.Open("/home/vadim/Desktop/myFiles/copy.odt")
	// check(err)

	// defer f.Close()

	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Printf("%d byte %s\n", n1, string(b1[0:n1]))

	// o2, err := f.Seek(6, 0)
	// check(err)
	// b2 := make([]byte, 2)
	// n2, err := f.Read(b2)
	// check(err)
	// fmt.Printf("%d bytes @ %d", n2, o2)
	// fmt.Printf("%v\n", b2[:n2])

	// o3, err := f.Seek(6, 0)
	// check(err)
	// b3 := make([]byte, 5)
	// n3, err := io.ReadAtLeast(f, b3, 5)
	// check(err)
	// fmt.Println(n3, o3, string(b3))

	// _, err = f.Seek(0, 0)
	// check(err)

	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(100)
	// check(err)
	// fmt.Printf("5 bytes: %s\n", string(b4))

	// f, err := os.Create("/home/vadim/Desktop/myFiles/copy.odt")
	// check(err)

	// f, _ := os.Open("/home/vadim/Desktop/myFiles/copy.odt")
	// defer f.Close()

	// defer os.Remove(f.Name())

	// os.Remove("/home/vadim/Desktop/myFiles/copy.odt")

	// f := "/home/vadim/Desktop/myFiles/copy.odt"
	// d2 := "i can help resrtgfwdrfgyhuioigfdsdzxfjgjkkjmhgbfd\n"
	// n2, _ := f.WriteString(d2)
	// fmt.Sprintln(n2)

	// f.Sync()

	// d1 := []byte("heroes, wwwwwwwwwwwwwwwwwwwwwwwwwwwwww  google\n")
	// g := os.WriteFile("/home/vadim/Desktop/myFiles/copy.odt", d1, 0644)
	// fmt.Sprintln(g)

	// defer g.Close()

	// w := bufio.NewWriter(f)
	// n4, _ := w.WriteString("hello world")
	// fmt.Println(n4)
	// w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "errors:", err)
	// 	os.Exit(1)
	// }

	// err := os.Mkdir("testdir", 0750)
	// if err != nil && !os.IsExist(err) {
	// 	log.Fatal(err)
	// }
	// err = os.WriteFile("testdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
