package main

import "fmt"

/*
	Разработать программу, которая переворачивает порядок слов в строке.

	Пример: входная строка:

	«snow dog sun», выход: «sun dog snow».

	Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы,
	а выполнять операцию «на месте».
*/

func main() {
	fmt.Println(reverseWord("snow dog sun"))
}

func reverseWord(s string) string {
	runes := []rune(s)

	reverse(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}
