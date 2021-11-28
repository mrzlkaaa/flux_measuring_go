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

type FoilData struct {
	Id             int64
	Name           string
	Nucleus_number float64
	Foil_type      string
}

// type StoreFoilsStore struct {
// 	MultiStore FoilsStore
// }

func (s *FoilsStore) TableName() string {
	return "foils_store"
}

func (s *FoilData) TableName() string {
	return "foil_data"
}

func FoilsStoreConstructor() *FoilsStore {
	return &FoilsStore{}
}

func StoreFoilsStoreConstructor() *[]FoilsStore {
	return &[]FoilsStore{}
}
