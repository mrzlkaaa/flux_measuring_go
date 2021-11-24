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

func PopulateStruct(nuclide string, connection func(dbname string) *gorm.DB) *FoilsStore { //map[string]interface{} {
	s := Prerequisites()
	fmt.Println(s)
	db := connection("foilsstore")
	db.First(s, "nuclide = ?", nuclide)
	fmt.Println(s)
	return s
	// return QueryFormatting(*allSamples)
}

//todo create func to request all params
