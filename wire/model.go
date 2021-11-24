package wire

type SetTableName interface {
	TableName() string
}

type Experiment struct {
	ID      int64
	Name    string   `gorm:"default:newTestInsert"`
	Samples []Sample `gorm:"foreignKey:Exp_id"`
}

type Sample struct {
	Name     int64
	Activity float64
	Exp_id   int64
}

func (exp *Experiment) TableName() string {
	return "experiment"
}

func (smpl *Sample) TableName() string {
	return "sample"
}

// func (*[]Sample) TableName() string {
// 	return "sample"
// }

func ExperimentConstructor(id int64) *Experiment {
	return &Experiment{ID: id}
	// return &Experiment{Name: name}
}

func SampleConstructor() *Sample {
	return &Sample{}
}

func AllSampleConstructor() *[]Sample {
	return &[]Sample{}
}
