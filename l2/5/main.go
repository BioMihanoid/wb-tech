package main

/*
	Что выведет программа?
	Объяснить вывод программы.

	Выведет "error". err переменная интерфейс error,
	когда мы туда записываем результат test, то мы возвращаем
	конкретный тип и его значение типа nil, поэтому сам интерфейс не равен nil.
*/

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
