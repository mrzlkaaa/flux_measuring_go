package detector_params

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func Prerequisites() *FoilsStore {
	s := FoilsStoreConstructor()
	s.TableName()
	return s
}

func (s *FoilsStore) Populate(value, param string, db *gorm.DB) *FoilsStore {
	// db := connection("foilsstore")
	query := fmt.Sprintf("%v = ?", param)
	db.First(s, query, value)
	s.Abundance = s.Abundance / 100
	to_float, _ := strconv.ParseFloat(s.Release, 64)
	s.Release = strconv.FormatFloat(to_float/100, 'f', 3, 64)
	fmt.Println(s)
	return s
}

func PopulateByAll(db *gorm.DB) *[]FoilsStore {
	var ms []FoilsStore
	db.Find(&ms)
	fmt.Println(ms)
	return &ms
}
