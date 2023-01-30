// Go program to illustrate how ...

package gM

import "fmt"

func GoMySolution5() {

	m_a_p := map[int]string{90: "Dog", 91: "Cat", 92: "Cow", 93: "Bird", 94: "Rabbit"}
	delete(m_a_p, 90)
	for lll, yyy := range m_a_p {

		fmt.Println(lll, yyy)
	}
	fmt.Println(m_a_p[93])

	well, well1 := m_a_p[94]
	_, well2 := m_a_p[95]
	fmt.Printf("%s\n", well)
	fmt.Printf("%t\n", well1)
	fmt.Printf("%t\n", well2)

	MP := map[string]int{"Yellow": 1, "Orange": 2, "Purple": 3, "Green": 4, "BLack": 5}
	fmt.Println("\ninitial map :", MP)
	MP2 := MP
	MP2["Brown"] = 6
	MP2["Red"] = 7
	fmt.Println("\nNew map :", MP2)
	fmt.Println("\nOld map in a new cover :", MP)

	//Range tasks

	var arr = [3]bool{true, false, true}
	for index := range arr {
		fmt.Printf("\n index[%t] = %t \n ", index, arr)
	}
}
