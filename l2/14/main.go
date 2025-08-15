package main

import (
	"fmt"
	"time"
)

func merge(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	doneCh := make(chan interface{})
	go func() {
		defer close(doneCh)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			mid := len(channels) / 2
			select {
			case <-merge(channels[:mid]...):
			case <-merge(channels[mid:]...):
			}
		}
	}()

	return doneCh
}

func createCh(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-merge(
		createCh(time.Microsecond),
		createCh(time.Millisecond),
		createCh(time.Millisecond),
		createCh(time.Second),
		createCh(time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
