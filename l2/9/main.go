package main

import (
	"fmt"

	"wb-tech/l2/9/pack"
)

/*
	Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.
	Примеры работы функции:
	Вход: "a4bc2d5e"
	Выход: "aaaabccddddde"

	Вход: "abcd"
	Выход: "abcd" (нет цифр — ничего не меняется)

	Вход: "45"
	Выход: "" (некорректная строка, т.к. в строке только цифры — функция должна вернуть ошибку)

	Вход: ""
	Выход: "" (пустая строка -> пустая строка)

	Дополнительное задание
	Поддерживать escape-последовательности вида \:

	Вход: "qwe\4\5"
	Выход: "qwe45" (4 и 5 не трактуются как числа, т.к. экранированы)

	Вход: "qwe\45"
	Выход: "qwe44444" (\4 экранирует 4, поэтому распаковывается только 5)

	Требования к реализации
	Функция должна корректно обрабатывать ошибочные случаи (возвращать ошибку, например, через error),
	и проходить unit-тесты.
	Код должен быть статически анализируем (vet, golint).
*/

func main() {
	res, err := pack.UnpackingStr("a4bc2d5e")
	fmt.Println(res, err)

	res, err = pack.UnpackingStr("abcd")
	fmt.Println(res, err)

	res, err = pack.UnpackingStr("45")
	fmt.Println(res, err)

	res, err = pack.UnpackingStr("")
	fmt.Println(res, err)

	res, err = pack.UnpackingStr("qwe\\4\\5")
	fmt.Println(res, err)

	res, err = pack.UnpackingStr("qwe\\45")
	fmt.Println(res, err)
}
