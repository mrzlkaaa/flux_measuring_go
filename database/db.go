package database

import (
	"flux_meas/detector_params"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	Populate(value, param string) detector_params.FoilsStore
}

type storage struct {
	db *gorm.DB
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

func NewConnection(db *gorm.DB) Storage {
	return &storage{db: db}
}

func (s *storage) Populate(value, param string) detector_params.FoilsStore {
	var params detector_params.FoilsStore
	// db := connection("foilsstore")
	query := fmt.Sprintf("%v = ?", param)
	s.db.First(params, query, value)
	params.Abundance = params.Abundance / 100
	to_float, _ := strconv.ParseFloat(params.Release, 64)
	params.Release = strconv.FormatFloat(to_float/100, 'f', 3, 64)
	fmt.Println(params)
	return params
}
