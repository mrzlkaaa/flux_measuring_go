package detector_params

import (
	"fmt"

	"gorm.io/gorm"
)

func Prerequisites() *FoilsStore {
	s := FoilsStoreConstructor()
	s.TableName()
	return s
}

// func (s *FoilsStore) Populate(value, param string, connection func(dbname string) *gorm.DB) *FoilsStore {
// 	db := connection("foilsstore")
// 	query := fmt.Sprintf("%v = ?", param)
// 	db.First(s, query, value)
// 	fmt.Println(s)
// 	return s
// }

func (s *FoilsStore) Populate(value, param string, db *gorm.DB) *FoilsStore {
	// db := connection("foilsstore")
	query := fmt.Sprintf("%v = ?", param)
	db.First(s, query, value)
	fmt.Println(s)
	return s
}

// func PopulateByAll(connection func(dbname string) *gorm.DB) *[]FoilsStore {
// 	db := connection("foilsstore")
// 	ss := &[]FoilsStore{}
// 	db.Find(ss)
// 	fmt.Println(ss)
// 	return ss
// }

// func PopulateStruct(nuclide string, connection func(dbname string) *gorm.DB) *FoilsStore { //map[string]interface{} {
// 	s := Prerequisites()
// 	fmt.Println(s)
// 	db := connection("foilsstore")
// 	db.First(s, "nuclide = ?", nuclide)
// 	fmt.Println(s)
// 	return s
// 	// return QueryFormatting(*allSamples)
// }

// func PopulateByNuclide(nuclide string, connection func(dbname string) *gorm.DB) *FoilsStore { //map[string]interface{} {
// 	s := Prerequisites()
// 	fmt.Println(s)
// 	db := connection("foilsstore")
// 	db.First(s, "nuclide = ?", nuclide)
// 	fmt.Println(s)
// 	return s
// 	// return QueryFormatting(*allSamples)
// }

//todo create func to request all params
