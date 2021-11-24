package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=irt-t.ru user=postgres password=postgres dbname=%v port=1111 sslmode=disable", dbname)
	// dsn := "host=db user=postgres password=postgres dbname=Detectors port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection failed")
	}
	return db
}
