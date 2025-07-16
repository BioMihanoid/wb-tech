package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.

	Классические подходы: выход по условию, через канал уведомления, через контекст,
	прекращение работы runtime.Goexit() и др.

	Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
}

// пример остановки горутины по условию
func example1() {
	stop := false
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for !stop {
			fmt.Println("running...")
			time.Sleep(time.Second)
		}
		fmt.Println("goroutine stopped")
	}()

	time.Sleep(3 * time.Second)
	stop = true
	wg.Wait()
	fmt.Println("goroutine stopped by condition")
}

// пример остановки горутины через канал
func example2() {
	stopCh := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				fmt.Println("goroutine stopped")
				return
			default:
				fmt.Println("running...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	stopCh <- struct{}{}
	wg.Wait()
	fmt.Println("goroutine stopped by channel")
}

// пример остановки горутины через контекст
func example3() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped")
				return
			default:
				fmt.Println("running...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("goroutine stopped by context cancellation")
}

// пример остановки горутины с помощью runtime.Goexit
func example4() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if i == 2 {
				fmt.Println("Calling Goexit")
				runtime.Goexit()
			}
			fmt.Println("Working", i)
			time.Sleep(time.Second)
		}
		fmt.Println("This will never print")
	}()

	wg.Wait()
	fmt.Println("goroutine stopped by runtime.Goexit")
}

// пример остановки горутины с использованием канала для получения данных
func example5() {
	dataCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range dataCh {
			fmt.Println("Received:", val)
		}
	}()

	for i := 0; i < 10; i++ {
		dataCh <- i
	}

	close(dataCh)
	wg.Wait()
	fmt.Println("Channel closed, goroutine will stop")
}
