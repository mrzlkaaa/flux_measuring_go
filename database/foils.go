package database

import (
	"flux_meas/detecParams"
	"sort"
)

func (s *storage) PopulateAllByFoilType(value string) ([]detecParams.FoilData, error) {
	var ms []detecParams.FoilData
	var query []detecParams.FoilData = []detecParams.FoilData{{Foil_type: value}}
	status := s.db.Where(&query).Find(&ms)
	sort.Slice(ms, func(i, j int) bool { return ms[i].Name < ms[j].Name })
	return ms, status.Error
}
