package provinsi

type Service interface {
	// GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error)
	GetProvinsi(page int, pageSize int) ([]Provinsi, error)
	GetTotalProvinsiCount() (int64, error)
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

func (s *service) GetProvinsi(page int, pageSize int) ([]Provinsi, error) {

	provinsis, err := s.repository.GetProvinsi(page, pageSize)

	return provinsis, err
}

func (s *service) GetTotalProvinsiCount() (int64, error) {

	kabupatentotal, err := s.repository.GetTotalProvinsiCount()
	return kabupatentotal, err

}
