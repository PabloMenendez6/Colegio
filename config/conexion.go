package config

import (
	"database/sql"
)

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5433 user=postgres password=postgres dbname=colegio_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
