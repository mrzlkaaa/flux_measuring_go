package detector_params

type FoilsStore struct {
	Id            int64
	Nuclide       string
	Cross_section float64
	Abundance     float64
	Half_life     float64
	Energy        string
	Release       string
	Resonance     float64
	Endf_data     string
}

// type StoreFoilsStore struct {
// 	MultiStore FoilsStore
// }

func (s *FoilsStore) TableName() string {
	return "foils_store"
}

func FoilsStoreConstructor() *FoilsStore {
	return &FoilsStore{}
}

func StoreFoilsStoreConstructor() *[]FoilsStore {
	return &[]FoilsStore{}
}
