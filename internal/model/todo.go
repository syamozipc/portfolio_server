package model

import "time"

type Todo struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
