package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Реализовать безопасную для конкуренции запись данных в структуру map.

	Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

	Проверьте работу кода на гонки (util go run -race).
*/

type SyncMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func (s *SyncMap[K, V]) Set(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SyncMap[K, V]) Get(key K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func main() {
	m := NewSyncMap[int, int]()

	go func() {
		for i := 0; i < 10; i++ {
			m.Set(i, i*i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			m.Set(i, i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v, ok := m.Get(i)
			if ok {
				fmt.Println(v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v, ok := m.Get(i)
			if ok {
				fmt.Println(v)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
