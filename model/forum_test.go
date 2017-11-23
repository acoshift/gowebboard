package model_test

import (
	"testing"

	"github.com/acoshift/gowebboard/model"
)

func TestFindAllForum(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	db.Exec(`insert into forums (id, title) values (1, 't1')`)
	db.Exec(`insert into forums (id, title) values (2, 't2')`)
	db.Exec(`insert into forums (id, title) values (3, 't3')`)

	forums, err := model.FindAllForum(db)
	if err != nil {
		t.Fatalf("expected FindAllForum returns nil error; got %v", err)
	}
	if l := len(forums); l != 3 {
		t.Fatalf("expected FindAllForum returns 3 forums; got %d", l)
	}
}

func TestCreateForum(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	id, err := model.CreateForum(db, &model.Forum{
		Title: "test",
	})
	if err != nil {
		t.Fatalf("expected CreateForum return nil error; got %v", err)
	}
	if id <= 0 {
		t.Fatalf("expected CreateForum returns positive id; got %d", id)
	}

	var cnt int
	db.QueryRow(`select count(*) from forums`).Scan(&cnt)
	if cnt != 1 {
		t.Fatalf("expected CreateForum insert 1 forum; got %d", cnt)
	}
}
