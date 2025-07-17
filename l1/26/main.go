package main

import (
	"fmt"
	"unicode"
)

/*
	Разработать программу, которая проверяет, что все символы в строке встречаются один раз
	(т.е. строка состоит из уникальных символов).

	Вывод: true, если все символы уникальны, false, если есть повторения.
	Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

	Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

	Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

func main() {
	fmt.Println(checkUniq("abcd"))
	fmt.Println(checkUniq("abCdefAaf"))
	fmt.Println(checkUniq("aabcd"))
}

func checkUniq(s string) bool {
	m := make(map[rune]struct{})
	for _, v := range s {
		if _, ok := m[unicode.ToLower(v)]; !ok {
			m[unicode.ToLower(v)] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
