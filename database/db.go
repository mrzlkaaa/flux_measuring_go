package database

import (
	"flux_meas/detecParams"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	PopulateByParam(value, param string) (detecParams.FoilsStore, error)
	PopulateByAll() ([]detecParams.FoilsStore, error)
	PopulateAllByFoilType(value string) ([]detecParams.FoilData, error)
	PopulateByQueryId(id int64) (map[string]interface{}, error)
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

// func (s *storage) PopulateByAll() ([]detector_params.FoilsStore, error) {
// 	var all_params []detector_params.FoilsStore
// 	status := s.db.Find(&all_params)
// 	return all_params, status.Error
// }
