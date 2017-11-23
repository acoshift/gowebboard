package model_test

import (
	"database/sql"
	"io/ioutil"
	"testing"

	_ "github.com/lib/pq"
)

func prepareDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres@localhost:5432?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// gowebboard_test
	_, err = db.Exec(`drop database if exists gowebboard_test`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`create database gowebboard_test`)
	if err != nil {
		t.Fatal(err)
	}

	db.Close()

	db, err = sql.Open("postgres", "postgres://postgres@localhost:5432/gowebboard_test?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	schema, err := ioutil.ReadFile("../model.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		t.Fatal(err)
	}

	return db
}
