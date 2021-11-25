package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Db *gorm.DB
}

func DbConnection(dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=irt-t.ru user=postgres password=postgres dbname=%v port=1111 sslmode=disable", dbname)
	// dsn := "host=db user=postgres password=postgres dbname=Detectors port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection failed")
	}
	return conn
}

func NewConnection() *Storage {
	// instance1 := DbConnection("Detectors")
	// db1 := &Storage{db: instance1}
	instance2 := DbConnection("foilsstore")
	// db2 := &Storage{db: instance2}
	return &Storage{Db: instance2}
}
