package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	dsn := "host=irt-t.ru user=postgres password=postgres dbname=Detectors port=1111 sslmode=disable"
	// dsn := "host=db user=postgres password=postgres dbname=Detectors port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection failed")
	}
	return db
}