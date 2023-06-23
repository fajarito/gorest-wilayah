package kecamatan

type Service interface {
	// GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error)
	GetKecamatan(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, page int, pageSize int) ([]Kecamatan, error)
	GetTotalKecamatanCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) (int64, error)
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

func (s *service) GetKecamatan(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, page int, pageSize int) ([]Kecamatan, error) {

	kecamatans, err := s.repository.GetKecamatan(KodeDepdagriProvinsi, KodeDepdagriKabupaten, page, pageSize)

	return kecamatans, err
}

func (s *service) GetTotalKecamatanCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string) (int64, error) {

	kecamatantotal, err := s.repository.GetTotalKecamatanCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten)
	return kecamatantotal, err

}
