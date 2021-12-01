package detecParams

type ParamsService interface {
	PopulateByParam(value, param string) (FoilsStore, error)
	PopulateByAll() ([]FoilsStore, error)
	PopulateAllByFoilType(value string) ([]FoilData, error)
}

type ParamsRepo interface {
	PopulateByParam(value, param string) (FoilsStore, error)
	PopulateByAll() ([]FoilsStore, error)
	PopulateAllByFoilType(value string) ([]FoilData, error)
}

type paramsService struct {
	storage ParamsRepo
}

func NewParamsService(db ParamsRepo) ParamsService {
	return &paramsService{storage: db}
}

func (s *paramsService) PopulateByParam(value, param string) (FoilsStore, error) {
	pop, err := s.storage.PopulateByParam(value, param)
	return pop, err
}

func (s *paramsService) PopulateByAll() ([]FoilsStore, error) {
	pop_all, err := s.storage.PopulateByAll()
	return pop_all, err
}

func (s *paramsService) PopulateAllByFoilType(value string) ([]FoilData, error) {
	pop_all, err := s.storage.PopulateAllByFoilType(value)
	return pop_all, err
}
