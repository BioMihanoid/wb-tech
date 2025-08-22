package model

import "time"

type Event struct {
	ID     int64     `json:"id"`
	UserID int64     `json:"user_id"`
	Date   time.Time `json:"date"`
	Text   string    `json:"event"`
}
