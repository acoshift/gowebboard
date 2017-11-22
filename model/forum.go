package model

import (
	"database/sql"
	"time"
)

type Forum struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindAllForum(db *sql.DB) ([]*Forum, error) {
	rows, err := db.Query(`
        select
            id, title, created_at, updated_at
        from forums
        order by created_at desc
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forums []*Forum
	for rows.Next() {
		var forum Forum
		err = rows.Scan(
			&forum.ID,
			&forum.Title,
			&forum.CreatedAt,
			&forum.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		forums = append(forums, &forum)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return forums, nil
}

func CreateForum(db *sql.DB, forum *Forum) (int, error) {
	var id int
	err := db.QueryRow(`
        insert into forums (
            title
        ) values (
            $1
        ) returning id
    `, forum.Title).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteForum(db *sql.DB, forumID int) error {
	_, err := db.Exec(`
        delete from forums
        where id = $1
    `, forumID)
	if err != nil {
		return err
	}
	return nil
}
