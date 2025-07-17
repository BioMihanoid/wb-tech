package main

import "fmt"

/*
	Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

	Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	for k, _ := range m {
		fmt.Print(k + " ")
	}
}
