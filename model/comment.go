package model

import "time"

type Comment struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
