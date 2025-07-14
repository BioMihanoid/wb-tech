package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

	Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func main() {
	t := 5
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <workerCount>")
		return
	}

	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil || workerCount <= 0 {
		fmt.Println("invalid worker count. Please provide a positive integer.")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	mainChannel := make(chan int)
	doneCh := make(chan struct{})

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-time.After(time.Duration(t) * time.Second)
		cancel()
		signalCh <- syscall.SIGTERM
		fmt.Println("time is up")
	}()

	for i := 0; i < workerCount; i++ {
		go worker(ctx, i, mainChannel, doneCh)
	}

	go func() {
		counter := 0
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Main goroutine exiting")
				close(mainChannel)
				return

			case mainChannel <- counter:
				counter++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	<-signalCh
	fmt.Println("Received shutdown signal")
	cancel()

	for i := 0; i < workerCount; i++ {
		<-doneCh
	}

	fmt.Println("Graceful shutdown complete")
}

func worker(ctx context.Context, id int, ch <-chan int, doneCh chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d exiting\n", id)
			doneCh <- struct{}{}
			return

		case value, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d channel closed, exiting\n", id)
				doneCh <- struct{}{}
				return
			}
			fmt.Printf("Worker %d received: %d\n", id, value)
		}
	}
}
