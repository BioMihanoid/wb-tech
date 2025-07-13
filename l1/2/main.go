package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10],
	и выведет результаты в stdout.

	Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

func main() {
	var a []int
	a = append(a, 2, 4, 6, 8, 10)
	wg := &sync.WaitGroup{}

	for _, v := range a {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v * v)
		}()
	}

	wg.Wait()
}
