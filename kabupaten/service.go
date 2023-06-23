package kabupaten

type Service interface {
	// GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error)
	GetKabupaten(KodeDepdagriProvinsi string, page int, pageSize int) ([]Kabupaten, error)
	GetTotalKabupatenCount(KodeDepdagriProvinsi string) (int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

// func (s *service) GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error) {

// 	kelurahan, err := s.repository.GetKelurahanDetail(IdKelurahan)

// 	return kelurahan, err
// }

func (s *service) GetKabupaten(KodeDepdagriProvinsi string, page int, pageSize int) ([]Kabupaten, error) {

	kecamatans, err := s.repository.GetKabupaten(KodeDepdagriProvinsi, page, pageSize)

	return kecamatans, err
}

func (s *service) GetTotalKabupatenCount(KodeDepdagriProvinsi string) (int64, error) {

	kabupatentotal, err := s.repository.GetTotalKabupatenCount(KodeDepdagriProvinsi)
	return kabupatentotal, err

}
