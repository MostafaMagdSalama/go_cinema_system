package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func Setup_DB(conn_string string) (*sql.DB, error) {
	fmt.Println("inside setup")

	db, err := sql.Open("postgres", conn_string)

	if err != nil {
		return nil, errors.Wrap(err, "connection with db")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "error in db ping")
	}

	return db, nil
}
