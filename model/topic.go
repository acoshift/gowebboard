package model

import "time"

type Topic struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
