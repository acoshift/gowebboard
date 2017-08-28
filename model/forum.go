package model

import "time"

type Forum struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
