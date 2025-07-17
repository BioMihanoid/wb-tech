package main

import "fmt"

/*
	Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

	Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:]))
	и уменьшить длину слайса на 1.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5}

	arr = remove(arr, 2)

	fmt.Println(arr)
}

func remove(arr []int, i int) []int {
	if i < 0 || i > len(arr) {
		return arr
	}

	copy(arr[i:], arr[i+1:])

	return arr[:len(arr)-1]
}
