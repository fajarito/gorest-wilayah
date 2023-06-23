package kabupaten

import (
	"gorm.io/gorm"
)

type Repository interface {
	// FindByID(nik string) (Stunting, error)
	// FindByPusHamil(nik string) (Stunting, error)
	// FindByStunting(nik string) (Stunting, error)
	// FindByBaduta(nik string) (Stunting, error)
	//KrsByKec(KodeDepdagriProvinsi string, KodeDepdagriKabupaten string, KodeDepdagriKecamatan string) ([]Krk, error)
	// GetKelurahanDetail(IdKelurahan int) ([]Kelurahan, error)
	GetKabupaten(KodeDepdagriProvinsi string, page int, pageSize int) ([]Kabupaten, error)
	GetTotalKabupatenCount(KodeDepdagriProvinsi string) (int64, error)

	// FindByKec(ProvinsiID int, KabupatenID int, KecamatanID int) ([]Krs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) GetKelurahanDetail(IdKelurahan string) (Kelurahan, error) {
// 	var kelurahan Kelurahan
// 	// err := r.db.Find(&kelurahan, TempatPelayananKBID).Error
// 	err := r.db.Where("id_kelurahan = ?", IdKelurahan).Find(&kelurahan).Error
// 	return kelurahan, err
// }

func (r *repository) GetKabupaten(KodeDepdagriProvinsi string, page int, pageSize int) ([]Kabupaten, error) {
	var kabupatens []Kabupaten

	// err := r.db.Find(&faskeses, ProvinsiID).Error
	offset := (page - 1) * pageSize
	err := r.db.Where("kode_depdagri_provinsi = ?", KodeDepdagriProvinsi).Offset(offset).Limit(pageSize).Find(&kabupatens).Error

	if err != nil {
		return nil, err
	}

	return kabupatens, nil
}

func (r *repository) GetTotalKabupatenCount(KodeDepdagriProvinsi string) (int64, error) {
	var count int64
	if err := r.db.Model(&Kabupaten{}).Where("kode_depdagri_provinsi = ? ", KodeDepdagriProvinsi).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
