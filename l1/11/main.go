package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы,
	присутствующие и в первом, и во втором.

	Пример:
	A = {1,2,3}
	B = {2,3,4}
	Пересечение = {2,3}
*/

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	m := make(map[int]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			fmt.Println(v)
		}
	}
}
