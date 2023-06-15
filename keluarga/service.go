package keluarga

type Service interface {
	GetKeluarga(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, page int, pageSize int) ([]Keluarga, error)
	GetTotalKeluargaCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) (int64, error)
	GetKeluargaBeresiko(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int, page int, pageSize int) ([]Keluarga, error)
	GetTotalKeluargaBeresikoCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int) (int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetKeluarga(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, page int, pageSize int) ([]Keluarga, error) {

	keluargas, err := s.repository.GetKeluarga(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan, page, pageSize)

	return keluargas, err
}

func (s *service) GetTotalKeluargaCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string) (int64, error) {

	keluargatotal, err := s.repository.GetTotalKeluargaCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan)
	return keluargatotal, err

}

func (s *service) GetKeluargaBeresiko(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int, page int, pageSize int) ([]Keluarga, error) {

	keluargas, err := s.repository.GetKeluargaBeresiko(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan, filter2, page, pageSize)

	return keluargas, err
}

func (s *service) GetTotalKeluargaBeresikoCount(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string, KodeDepdagriKelurahan string, filter2 int) (int64, error) {

	keluargatotal, err := s.repository.GetTotalKeluargaBeresikoCount(KodeDepdagriProvinsi, KodeDepdagriKabupaten, KodeDepdagriKecamatan, KodeDepdagriKelurahan, filter2)
	return keluargatotal, err

}
