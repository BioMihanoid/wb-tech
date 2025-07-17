package main

import "fmt"

/*
	Разработать программу, которая переворачивает подаваемую на вход строку.

	Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

	Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
	то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).
*/

func main() {
	s := "главрыба"
	fmt.Println(reverseStr(s))
}

func reverseStr(s string) string {
	res := []rune(s)
	l := 0
	r := len(res) - 1

	for l <= r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return string(res)
}
