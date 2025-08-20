package service

import (
	"sync"

	"wb-tech/l2/18/internal/model"
)

type Service struct {
	mu   sync.RWMutex
	data map[int64][]model.Event
}

func NewService() *Service {
	return &Service{
		data: make(map[int64][]model.Event),
	}
}
