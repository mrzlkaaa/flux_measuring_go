package database

import (
	"flux_meas/detecParams"
	"fmt"
	"strconv"
)

func (s *storage) PopulateByParam(value, param string) (detecParams.FoilsStore, error) {
	var params detecParams.FoilsStore
	query := fmt.Sprintf("%v = ?", param)
	status := s.db.First(&params, query, value)
	params.Abundance = params.Abundance / 100
	to_float, _ := strconv.ParseFloat(params.Release, 64)
	params.Release = strconv.FormatFloat(to_float/100, 'f', 3, 64)
	fmt.Println(params)
	return params, status.Error
}

func (s *storage) PopulateByAll() ([]detecParams.FoilsStore, error) {
	var all_params []detecParams.FoilsStore
	status := s.db.Find(&all_params)
	return all_params, status.Error
}
