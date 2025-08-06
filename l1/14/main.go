package main

import "fmt"

/*
	Разработать программу, которая в runtime способна определить тип переменной,
	переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

	Подсказка: оператор типа switch v.(type) поможет в решении.
*/

func main() {
	detectType(123)
	detectType("hello")
	detectType(true)
	detectType(make(chan any))
	detectType(3.14)
}

func detectType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("this is string")
	case int:
		fmt.Println("this is int")
	case bool:
		fmt.Println("this is bool")
	case chan any:
		fmt.Println("this is chan any")
	default:
		fmt.Println("unknown type")
	}
}
