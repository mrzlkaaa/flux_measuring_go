package wire

import (
	"fmt"

	"gorm.io/gorm"
)

func QueryFormatting(sp []Sample) map[string]interface{} {
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

func RenameTable() {
	var smpl SetTableName = SampleConstructor()
	smpl.TableName()
}

func PopulateStruct(id int64, connection func(dbname string) *gorm.DB) map[string]interface{} {
	RenameTable()
	db := connection("Detectors")
	allSamples := AllSampleConstructor()
	db.Where(&[]Sample{{Exp_id: id}}).Find(allSamples)
	return QueryFormatting(*allSamples)
}
