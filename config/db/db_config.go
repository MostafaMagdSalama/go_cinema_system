package config

import (
	"database/sql"

	"log"
	"github.com/pkg/errors"
	_"github.com/lib/pq"
	
)

func Setup_DB(conn_string string) (*sql.DB, error) {
	log.Print("inside setup")

	db, err := sql.Open("postgres", conn_string)

	if err != nil {
		return nil, errors.Wrap(err, "connection with db ")

	}


	if err := db.Ping(); err != nil {
		errors.Wrap(err, "error in db ping")
	}

	return db, nil
}
