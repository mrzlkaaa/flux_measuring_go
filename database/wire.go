package database

import (
	"flux_meas/wire"
	"fmt"
)

func (s *storage) PopulateByQueryId(id int64) (map[string]interface{}, error) {
	var samples []wire.Sample
	var query []wire.Sample = []wire.Sample{{Exp_id: id}}
	status := s.db.Where(query).Find(&samples)
	formatted := QueryFormatting(samples)
	return formatted, status.Error
}

func QueryFormatting(sp []wire.Sample) map[string]interface{} {
	var sliceName []int64
	var sliceAct []float64
	for _, v := range sp {
		sliceName = append(sliceName, v.Name)
		sliceAct = append(sliceAct, v.Activity)
	}
	mapping := map[string]interface{}{
		"Name":     sliceName,
		"Activity": sliceAct,
	}
	fmt.Println(mapping)
	return mapping
}
