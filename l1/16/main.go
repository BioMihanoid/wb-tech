package main

import "fmt"

/*
	Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

	Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
	Для выбора опорного элемента можно взять середину или первый элемент.
*/

func main() {
	arr := []int{3, 2, 5, 1, 2, 4, 6, 7, 2, 3, 9, 3, 4, 1, 7, 3, 6}
	arr = quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, equal, greater []int

	for _, v := range arr {
		switch {
		case v < pivot:
			less = append(less, v)
		case v == pivot:
			equal = append(equal, v)
		case v > pivot:
			greater = append(greater, v)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(greater)...)
}
