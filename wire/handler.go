package wire

type WireService interface {
	PopulateByQueryId(id int64) (map[string]interface{}, error)
}

type WireRepo interface {
	PopulateByQueryId(id int64) (map[string]interface{}, error)
}

type wireService struct {
	storage WireRepo
}

// func RenameTable() {
// 	var smpl SetTableName = SampleConstructor()
// 	smpl.TableName()
// }

func NewWireService(db WireRepo) WireService {
	return &wireService{storage: db}
}

func (s *wireService) PopulateByQueryId(id int64) (map[string]interface{}, error) {
	get_data, err := s.storage.PopulateByQueryId(id)
	return get_data, err
}
