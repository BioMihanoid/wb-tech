package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b,
	значения которых > 2^20 (больше 1 миллион).

	Комментарий: в Go тип int справится с такими числами,
	но обратите внимание на возможное переполнение для ещё больших значений.
	Для очень больших чисел можно использовать math/big.
*/

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("10485760", 10)
	b.SetString("20971520", 10)

	sum := big.NewInt(0).Add(a, b)
	sub := big.NewInt(0).Sub(a, b)
	mul := big.NewInt(0).Mul(a, b)
	div := big.NewInt(0).Div(b, a)

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("Сумма:        %v\n", sum)
	fmt.Printf("Разность:     %v\n", sub)
	fmt.Printf("Произведение: %v\n", mul)
	fmt.Printf("Частное:      %v\n", div)
}
