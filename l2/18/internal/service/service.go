package service

import (
	"errors"
	"sync"
	"time"

	"wb-tech/l2/18/internal/model"
)

type Service struct {
	mu     sync.RWMutex
	data   map[int64][]model.Event
	nextID int64
}

func NewService() *Service {
	return &Service{
		data: make(map[int64][]model.Event),
	}
}

func (s *Service) CreateEvent(userID int64, date time.Time, text string) (model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range s.data[userID] {
		if v.Date == date && v.Text == text {
			return model.Event{}, errors.New("event already exists")
		}
	}

	ev := model.Event{
		ID:     s.nextID,
		UserID: userID,
		Date:   date,
		Text:   text,
	}
	s.nextID++

	s.data[userID] = append(s.data[userID], ev)
	return ev, nil
}

func (s *Service) UpdateEvent(userID, id int64, date time.Time, text string) (model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	events := s.data[userID]
	for i, v := range events {
		if v.ID == id {
			ev := model.Event{
				ID:     id,
				UserID: userID,
				Date:   date,
				Text:   text,
			}
			events[i] = ev
			s.data[userID] = events
			return ev, nil
		}
	}

	return model.Event{}, errors.New("event does not exist")
}

func (s *Service) DeleteEvent(userID, id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	events := s.data[userID]
	for i, v := range events {
		if v.ID == id {
			s.data[userID] = append(events[:i], events[i+1:]...)
			return nil
		}
	}

	return errors.New("event does not exist")
}

func (s *Service) EventsForDay(userID int64, date time.Time) ([]model.Event, error) {
	var evs []model.Event

	s.mu.RLock()

	for _, v := range s.data[userID] {
		if v.Date.Year() == date.Year() && v.Date.YearDay() == date.YearDay() {
			evs = append(evs, v)
		}
	}

	s.mu.RUnlock()

	if len(evs) == 0 {
		return nil, errors.New("no events found")
	}

	return evs, nil
}

func (s *Service) EventsForWeek(userID int64, date time.Time) ([]model.Event, error) {
	var evs []model.Event

	s.mu.RLock()

	year, week := date.ISOWeek()
	for _, v := range s.data[userID] {
		y, w := v.Date.ISOWeek()
		if y == year && w == week {
			evs = append(evs, v)
		}
	}

	s.mu.RUnlock()

	if len(evs) == 0 {
		return nil, errors.New("no events found")
	}

	return evs, nil
}

func (s *Service) EventsForMonth(userID int64, date time.Time) ([]model.Event, error) {
	var evs []model.Event

	s.mu.RLock()

	year, month := date.Year(), date.Month()
	for _, v := range s.data[userID] {
		if v.Date.Year() == year && v.Date.Month() == month {
			evs = append(evs, v)
		}
	}

	s.mu.RUnlock()

	if len(evs) == 0 {
		return nil, errors.New("no events found")
	}

	return evs, nil
}
