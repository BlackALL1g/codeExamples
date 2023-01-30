// Последовательно идущие единицы
// Требуется найти в бинарном векторе самую длинную последовательность единиц
// и вывести её длину.

// Желательно получить решение, работающее за линейное время
// и при этом проходящее по входному массиву только один раз.

// Формат ввода
// Первая строка входного файла содержит одно число n, n ≤ 10000.
//  Каждая из следующих n строк содержит ровно одно число — очередной элемент массива.

// Формат вывода
// Выходной файл должен содержать единственное число
//
//	— длину самой длинной последовательности единиц во входном массиве.

// Пример

// Ввод	Вывод
// 5		1
// 1
// 0
// 1
// 0
// 1

package main

import "fmt"

func consecutiveUnits() int {
	var lim int
	var b []int
	fmt.Scanln(&lim)
	var counter int

	for i := 0; i < lim; i++ {
		fmt.Scan(&counter)
		b = append(b, counter)
	}

	var counter2, storage int

	for _, v := range b {

		if v == 1 {
			counter2++
		} else {
			counter2 = 0
		}
		if counter2 > storage {
			storage = counter2
		}
	}

	return storage

}

func main() {
	fmt.Println(consecutiveUnits())
}
