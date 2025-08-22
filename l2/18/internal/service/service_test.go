package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"wb-tech/l2/18/internal/service"
)

func TestCreateEvent(t *testing.T) {
	s := service.NewService()
	date := time.Now()

	ev, err := s.CreateEvent(1, date, "test event")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), ev.ID)
	assert.Equal(t, int64(1), ev.UserID)
	assert.Equal(t, "test event", ev.Text)

	_, err = s.CreateEvent(1, date, "test event")
	assert.Error(t, err)
}

func TestUpdateEvent(t *testing.T) {
	s := service.NewService()
	date := time.Now()

	ev, _ := s.CreateEvent(1, date, "old event")

	updated, err := s.UpdateEvent(1, ev.ID, date.Add(time.Hour), "new event")
	assert.NoError(t, err)
	assert.Equal(t, "new event", updated.Text)

	_, err = s.UpdateEvent(1, 999, date, "fail")
	assert.Error(t, err)
}

func TestDeleteEvent(t *testing.T) {
	s := service.NewService()
	date := time.Now()

	ev, _ := s.CreateEvent(1, date, "to delete")

	err := s.DeleteEvent(1, ev.ID)
	assert.NoError(t, err)

	err = s.DeleteEvent(1, ev.ID)
	assert.Error(t, err)
}

func TestEventsForDayWeekMonth(t *testing.T) {
	s := service.NewService()
	baseDate := time.Date(2025, 8, 22, 10, 0, 0, 0, time.UTC)

	ev1, _ := s.CreateEvent(1, baseDate, "today event")
	ev2, _ := s.CreateEvent(1, baseDate.AddDate(0, 0, 2), "same week event")
	ev3, _ := s.CreateEvent(1, baseDate.AddDate(0, 1, 0), "next month event")

	evs, err := s.EventsForDay(1, baseDate)
	assert.NoError(t, err)
	assert.Len(t, evs, 1)
	assert.Equal(t, ev1.ID, evs[0].ID)

	evs, err = s.EventsForWeek(1, baseDate)
	assert.NoError(t, err)
	assert.Len(t, evs, 2)
	assert.ElementsMatch(t, []int64{ev1.ID, ev2.ID}, []int64{evs[0].ID, evs[1].ID})

	evs, err = s.EventsForMonth(1, baseDate)
	assert.NoError(t, err)
	assert.Len(t, evs, 2)
	assert.ElementsMatch(t, []int64{ev1.ID, ev2.ID}, []int64{evs[0].ID, evs[1].ID})

	evs, err = s.EventsForMonth(1, ev3.Date)
	assert.NoError(t, err)
	assert.Len(t, evs, 1)
	assert.Equal(t, ev3.ID, evs[0].ID)
}
