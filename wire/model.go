package wire

type SetTableName interface {
	TableName() string
}

type Experiment struct {
	ID      int64
	Name    string
	Samples []Sample
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
