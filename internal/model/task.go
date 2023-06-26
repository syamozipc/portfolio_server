package model

import "time"

type Task struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
