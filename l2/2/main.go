package main

import "fmt"

/*
	Что выведет программа?

	Объяснить порядок выполнения defer функций и итоговый вывод.

	В случае test она выведет 2, так как мы возвращаемаем именованную переменную,
	то после возврата она ещё может изменится дефером. В случае anotherTest мы
	получим ответ 1, так дефер будет работать уже со своей локальной копией переменной.
*/

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
