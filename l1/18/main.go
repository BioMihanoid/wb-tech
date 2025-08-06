package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
	По завершению программы структура должна выводить итоговое значение счётчика.

	Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

func main() {
	var count atomic.Int32
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			count.Add(1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			count.Add(1)
		}
	}()

	wg.Wait()
	fmt.Printf("count = %d", count.Load())
}
