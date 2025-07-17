package main

import "fmt"

/*
	Реализовать алгоритм бинарного поиска встроенными методами языка.
	Функция должна принимать отсортированный слайс и искомый элемент,
	возвращать индекс элемента или -1, если элемент не найден.

	Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.
*/

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println(binSearch(arr, 7))
	fmt.Println(binSearch(arr, 2))
}

func binSearch(arr []int, v int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) / 2

		if arr[mid] == v {
			return mid
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
