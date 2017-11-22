package model

import (
	"database/sql"
	"time"

	"github.com/acoshift/gowebboard/api"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindUserByUsername(db *sql.DB, username string) (*User, error) {
	var user User
	err := db.QueryRow(`
        select
            id, username, password, created_at, updated_at
        from users
        where username = $1
    `, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, api.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
