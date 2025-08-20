package service

import (
	"errors"
	"sync"
	"time"

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

func (s *Service) CreateEvent(userID int64, date time.Time, text string) (model.Event, error) {
	return model.Event{}, errors.New("not implemented")
}

func (s *Service) UpdateEvent(userID int64, date time.Time, text string) (model.Event, error) {
	return model.Event{}, errors.New("not implemented")
}

func (s *Service) DeleteEvent(userID, id int64) error {
	return errors.New("not implemented")
}

func (s *Service) EventsForDay(userID int64, date time.Time) ([]model.Event, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) EventsForWeek(userID int64, date time.Time) ([]model.Event, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) EventsForMonth(userID int64, date time.Time) ([]model.Event, error) {
	return nil, errors.New("not implemented")
}
