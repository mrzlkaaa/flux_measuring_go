package detector_params

import (
	"fmt"
	"sort"

	"gorm.io/gorm"
)

func Prerequisites() *FoilsStore {
	s := FoilsStoreConstructor()
	s.TableName()
	return s
}

type ParamsService interface {
	Populate(value, param string) error
}

type ParamsRepo interface {
	Populate(value, param string) FoilsStore
}

type paramsService struct {
	storage ParamsRepo
}

func NewParamsService(db ParamsRepo) ParamsService {
	return &paramsService{storage: db}
}

func (s *paramsService) Populate(value, param string) error {
	s.storage.Populate(value, param)
	return nil
}

func PopulateByAll(db *gorm.DB) *[]FoilsStore {
	var ms []FoilsStore
	db.Find(&ms)
	fmt.Println(ms)
	return &ms
}

func PopulateByFoilNames(value string, db *gorm.DB) *[]FoilData {
	var ms []FoilData
	db.Where(&[]FoilData{{Foil_type: value}}).Find(&ms)
	sort.Slice(ms, func(i, j int) bool { return ms[i].Name < ms[j].Name })

	return &ms
}
