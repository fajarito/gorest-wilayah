package kelurahan

type Service interface {
	// GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error)
	GetKelurahan(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, page int, pageSize int) ([]Kelurahan, error)
	GetTotalKelurahanCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) (int64, error)
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

func (s *service) GetKelurahan(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, page int, pageSize int) ([]Kelurahan, error) {

	kelurahans, err := s.repository.GetKelurahan(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, page, pageSize)

	return kelurahans, err
}

func (s *service) GetTotalKelurahanCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) (int64, error) {

	keluargatotal, err := s.repository.GetTotalKelurahanCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan)
	return keluargatotal, err

}
