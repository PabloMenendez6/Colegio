package config

import (
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	dns := "host=localhost port=5433 user=postgres password=postgres dbname=colegio_db sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
