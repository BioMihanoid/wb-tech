package main

import (
	"fmt"
)

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива,
	во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout.
	То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
	Убедитесь, что чтение из второго канала корректно завершается.
*/

func main() {
	in := make(chan int)
	out := make(chan int)

	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}

	go func() {
		for _, v := range arr {
			in <- v
		}
		close(in)
	}()

	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}

}
